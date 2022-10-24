const popup = document.querySelector('.popup-login');
const login = document.getElementById('loginResult');
popup.classList.add('visually-hidden');
console.log(login.textContent);

if (login.textContent === '1') {
  popup.classList.remove('visually-hidden');
  window.setTimeout(addHidden, 5000);
  function addHidden() {
    popup.classList.add('visually-hidden');
  }
}
