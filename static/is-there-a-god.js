// Initialise question field
let questions = [];

// Create 'form' to link html to js, selects the html form
const form = document.querySelector('.form');

// Create the list of questions and links to html
const questionList = document.querySelector('#question-list');

// Function to render the list of questions
function renderQuestions() {
  // Clear the existing list
  questionList.innerHTML = '';

  // Create a new list item for each question
  questions.forEach((question) => {
    const listItem = document.createElement('li');
    const title = document.createElement('h3');
    const content = document.createElement('p');
    const answerForm = document.createElement('form');
    const answerInput = document.createElement('input');
    const answerButton = document.createElement('button');

    title.innerText = question.title;
    content.innerText = question.content;
    answerInput.setAttribute('type', 'text');
    answerInput.setAttribute('placeholder', 'Type your answer here:');
    answerButton.innerText = 'Submit';

    answerForm.addEventListener('submit', (event) => {
      event.preventDefault();
      const answer = answerInput.value.trim();
      if (answer) {
        question.answers.push({
          id: Date.now(),
          content: answer,
          votes: 0
        });
        answerInput.value = '';
        renderQuestions();
      }
    });

    listItem.appendChild(title);
    listItem.appendChild(content);

    const answerList = document.createElement('ul');
    answerList.setAttribute('class', 'answer-list');

    question.answers.forEach((answer) => {
      const answerItem = document.createElement('li');
      const answerContent = document.createElement('p');
      const voteForm = document.createElement('form');
      const voteButton = document.createElement('button');
      const voteCount = document.createElement('span');

      answerContent.innerText = answer.content;
      voteButton.innerText = 'Vote';
      voteCount.innerText = answer.votes.toString();

      voteForm.addEventListener('submit', (event) => {
        event.preventDefault();
        answer.votes += 1;
        voteCount.innerText = answer.votes.toString();
      });

      voteForm.appendChild(voteButton);
      voteForm.appendChild(voteCount);
      answerItem.appendChild(answerContent);
      answerItem.appendChild(voteForm);
      answerList.appendChild(answerItem);
    });

    listItem.appendChild(answerList);
    listItem.appendChild(answerForm);
    answerForm.appendChild(answerInput);
    answerForm.appendChild(answerButton);

    questionList.appendChild(listItem);
  });
}

// Function to add a new question to the questions array
function addQuestion() {
  const input = document.querySelector('#question');
  const title = input.value.trim();
  if (title) {
    const question = {
      id: Date.now(),
      title,
      content: '',
      votes: 0,
      answers: []
    };
    questions.push(question);
    input.value = '';
    renderQuestions();
  }
}

// Add event listener to submit button to add new question
const submitButton = document.querySelector('#submit');
submitButton.addEventListener('click', addQuestion);

// Get profile data from local storage, or create an empty object if none exists
let profileData = JSON.parse(localStorage.getItem('profileData')) || {};

// Create variables for profile form elements
const profileForm = document.querySelector('#profile-form');
const profileNameInput = document.querySelector('#name');
const profileEmailInput = document.querySelector('#email');
const profilePicInput = document.querySelector('#profile-pic');
const profilePicPreview = document.querySelector('.profile-pic');