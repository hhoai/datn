// Function to fetch language data
async function fetchLanguageData(lang) {
    const response = await fetch(`/main/languages/${lang}.json`);
    return response.json();
}

// Function to set the language preference
function setLanguagePreference(lang) {
    localStorage.setItem("language", lang);
    location.reload();
}

// Function to update content based on selected language
function updateContent(langData) {
    document.querySelectorAll("[data-i18n]").forEach((element) => {
        const key = element.getAttribute("data-i18n");

        if (element.tagName === "INPUT" && key === "placeholder_text") {
            // If the element is an input with placeholder_text attribute, set placeholder
            element.placeholder = langData[key];
        } else {
            // For other elements, set text content
            //element.textContent = langData[key];
            element.innerHTML = langData[key];
        }
    });
}

function updateImage(lang) {
    // Select the img element by its ID
    const imgElement = document.getElementById("img-header");

    switch (lang) {
        case "en":
            imgElement.src = "/images/great-britain.png";
            break;
        case "vn":
            imgElement.src = "/images/vietnam.png";
            break;
    }
}

// Function to change language
async function changeLanguage(lang) {
    await setLanguagePreference(lang);
    updateImage(lang);
    const langData = await fetchLanguageData(lang);
    updateContent(langData);

    //
    toggleArabicStylesheet(lang); // Toggle Arabic stylesheet
}

// Function to toggle Arabic stylesheet based on language selection
function toggleArabicStylesheet(lang) {
    const head = document.querySelector("head");
    const link = document.querySelector("#styles-link");

    if (link) {
        head.removeChild(link); // Remove the old stylesheet link
    }
    // else if (lang === "ar") {
    //   const newLink = document.createElement("link");
    //   newLink.id = "styles-link";
    //   newLink.rel = "stylesheet";
    //   newLink.href = "./assets/css/style-ar.css"; // Path to Arabic stylesheet
    //   head.appendChild(newLink);
    // }
}

// Call updateContent() on page load
window.addEventListener("DOMContentLoaded", async () => {
    const userPreferredLanguage = localStorage.getItem("language") || "en";
    const langData = await fetchLanguageData(userPreferredLanguage);
    updateImage(userPreferredLanguage);
    updateContent(langData);
    toggleArabicStylesheet(userPreferredLanguage);
});

function showModalSuccess(message){
    $.toast({
        heading:  'Successfully!',
        text: message,
        position: 'top-right',
        loaderBg: '#ff6849',
        icon: 'success',
        hideAfter: 2000,
        stack: 6
    });
}

function showModalError(message){
    $.toast({
        heading:  'Warning!',
        text: message,
        position: 'top-right',
        loaderBg: '#ff6849',
        icon: 'error',
        hideAfter: 2000,
        stack: 6
    });
}