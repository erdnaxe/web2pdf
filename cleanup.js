/*
 * Cleanup script: fix elements that are problematic when printing a web page
 */

searchLink = document.querySelector("link[rel='search']")
if (searchLink && searchLink.title === "Medium") {
    // Hide Medium privacy policy banner
    document.evaluate("//a[contains(., 'Privacy Policy')]/../../../..", document).iterateNext().hidden = true
}    
