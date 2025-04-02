package main

import (
	"crypto/tls"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"lms/handles"
	"lms/repo"
	"lms/routers"
	"lms/utilS"
	"os"
	"strings"
)

func loadProtocolHttps(app *fiber.App) {
	certCrtPaths := strings.Split(os.Getenv("CERT_SRV"), ",")
	certKeyPaths := strings.Split(os.Getenv("KEY_SRV"), ",")
	domains := strings.Split(os.Getenv("DOMAIN"), ",")

	if len(certCrtPaths) != len(certKeyPaths) || len(certCrtPaths) != len(domains) {
		repo.OutPutDebugError("Mismatched lengths of certificate paths, key paths, and domains")
	}

	certificates := make(map[string]*tls.Certificate)

	for i, domain := range domains {
		cert, err := tls.LoadX509KeyPair(certCrtPaths[i], certKeyPaths[i])
		if err != nil {
			repo.OutPutDebugError("Failed to load certificate for" + domain + err.Error())
		}
		certificates[domain] = &cert
	}

	tlsConfig := &tls.Config{
		GetCertificate: func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
			if cert, exists := certificates[info.ServerName]; exists {
				return cert, nil
			}
			return nil, nil
		},
	}

	ln, err := tls.Listen("tcp", ":"+os.Getenv("PORT_SRV"), tlsConfig)
	if err != nil {
		repo.OutPutDebugError("failed to start TLS listener: " + err.Error())
	}

	if err := app.Listener(ln); err != nil {
		repo.OutPutDebugError(err.Error())
	}
}

func loadProtocolHttp(app *fiber.App) {
	repo.OutPutDebugError(":" + os.Getenv("PORT_SRV"))
	if err := app.Listen(":" + os.Getenv("PORT_SRV")); err != nil {

		repo.OutPutDebugError(err.Error())
	}
}

func main() {
	//if err := repo.MigrateDB(); err != nil {
	//	repo.OutPutDebugError(err.Error())
	//}

	engine := html.New("./views", ".html")

	engine.AddFuncMap(fiber.Map{
		"CheckPermission":      handlers.CheckPermission,
		"CheckActivateAccount": handlers.CheckActivateAccount,
		"CheckTypeUser":        handlers.CheckTypeUser,
	})

	app := fiber.New(fiber.Config{
		Views:             engine,
		PassLocalsToViews: true,
		StrictRouting:     true,
		CaseSensitive:     true,
		UnescapePath:      true,
		BodyLimit:         50 * 1024 * 1024, // 50MB
	})

	app.Use(func(c *fiber.Ctx) error {
		if strings.Contains(c.Path(), "..") ||
			strings.Contains(strings.ToLower(c.Path()), "script") ||
			strings.Contains(c.Path(), "(") ||
			strings.Contains(c.Path(), ";") ||
			strings.Contains(c.Path(), "--") {
			repo.OutPutDebugError(c.IP() + c.Path())
			return c.Redirect("/403")
		}
		return c.Next()
	})

	app.Static("/", "./public")
	app.Static("", "./documents")

	routers.Router(app)
	routers.API(app)

	utilS.DomainEmail = utilS.GetDomainEmail()
	if os.Getenv("PROTOCOL") == "https" {
		loadProtocolHttps(app)
	} else {
		loadProtocolHttp(app)
	}
}
