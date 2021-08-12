/*
 * Cleanup script: fix elements that are problematic when printing a web page
 */

// Convert display fixed to display static
// This removes cookie banners that are hidding content
[...document.body.getElementsByTagName("*")].filter(
    x => getComputedStyle(x, null).getPropertyValue("position") === "fixed"
).forEach(e => {
    e.style.position = "static";
});
