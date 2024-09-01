class RandomWords extends EventTarget {
  wordsArray = [];

  constructor(wordsArray) {
    super();

    this.wordsArray = wordsArray;
    this.value = wordsArray[0];
  }

  // valuechange event will be triggered where the details is the value of the object.
  #emitChangeEvent() {
    this.dispatchEvent(new CustomEvent("valuechange", { detail: this.value }));
  }

  generate() {
    const randomIndex = Math.floor(Math.random() * this.wordsArray.length);
    this.value = this.wordsArray[randomIndex];

    this.#emitChangeEvent();
  }
}

const initialValue = ["Hello", "Asteroids", "Human", "Universe"];
const randomWorder = new RandomWords(initialValue);

document.querySelector("#currentWord").textContent = initialValue[0];

randomWorder.addEventListener("valuechange", (event) => {
  document.querySelector("#currentWord").textContent = event.detail;
});

document.querySelector("#generate").addEventListener("click", () => {
  randomWorder.generate();
});
