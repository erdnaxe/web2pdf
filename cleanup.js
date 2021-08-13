// Copyright (c) 2021 Alexandre Iooss
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// Cleanup script: fix elements that are problematic when printing a web page

async () => {
    // Convert display fixed to display static
    // This removes cookie banners that are hidding content
    [...document.body.getElementsByTagName("*")].filter(
        x => getComputedStyle(x, null).getPropertyValue("position") === "fixed"
    ).forEach(e => e.style.position = "static")

    // Scroll the whole page to trigger lazy loading
    await new Promise((resolve) => {
        const timer = setInterval(() => {
            window.scrollBy(0, window.innerHeight)
            if (window.scrollY + window.innerHeight >= document.body.scrollHeight) {
                clearInterval(timer)
                resolve()
            }
        }, 100)
    })
}
