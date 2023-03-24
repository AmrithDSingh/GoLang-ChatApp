//Initialise question field
let questions = [];

//Create 'form' to link html to js, selects the html form
const form = document.querySelector('.form');

//Create the list of questions and links to html
const questionList = document.querySelector('#question-list');

// Function to render the list of questions
function renderQuestions() {
  // Clear the existing list
  questionList.innerHTML = '';
  
  // Create a new list item for each question
  questions.forEach((question) => {
    const listItem = document.createElement('li');
    listItem.innerText = question.title;
    questionList.appendChild(listItem);
  });
}

form.addEventListener('submit', (event) => {
  event.preventDefault();
  const input = document.querySelector('#question');
  const title = input.value.trim();
  if (title) {
    // Add the new question to the questions array
    const question = {
      id: Date.now(),
      title,
      content: '',
      votes: 0,
      answers: []
    };
    questions.push(question);
    // Clear the input field
    input.value = '';
    // Rerender the list of questions
    renderQuestions();
  }
});