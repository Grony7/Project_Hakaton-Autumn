const button = document.querySelector('.news__buttons-button--toggle');
const textFirst = 'Горячие';
const textSecond = 'Холодные';

button.onclick = (evt) => {
  evt.preventDefault();
  if (button.textContent === textFirst) {
    button.textContent = textSecond;
  } else {
    button.textContent = textFirst;
  }
};
