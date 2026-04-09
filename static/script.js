'use strict';

const form  = document.querySelector('.search-form');
const input = document.querySelector('.search-input');

if (form && input) {
    // Prevent full page reload — results update in place.
    form.addEventListener('submit', e => e.preventDefault());

    let timer;
    input.addEventListener('input', () => {
        clearTimeout(timer);
        timer = setTimeout(() => liveSearch(input.value.trim()), 300);
    });
}

async function liveSearch(query) {
    const url = query ? `/search?q=${encodeURIComponent(query)}` : '/';

    let html;
    try {
        html = await fetch(url).then(r => r.text());
    } catch {
        return; // silent fail on network error
    }

    const incoming = new DOMParser().parseFromString(html, 'text/html');
    const main     = document.querySelector('main');
    const newMain  = incoming.querySelector('main');

    if (main && newMain) main.innerHTML = newMain.innerHTML;
}
