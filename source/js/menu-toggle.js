const navMain = document.querySelector('.main-nav');
const navToggle = document.querySelector('.main-nav__toggle');

navMain.classList.remove('main-nav--nojs');

navToggle.onclick = (evt) => {
    evt.preventDefault();
    navMain.classList.toggle('main-nav--opened');
    navMain.classList.toggle('main-nav--closed');
};
