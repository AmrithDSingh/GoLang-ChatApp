//Initialise question field
let questions = [];
//Create 'form' to link html to js, selects the html form
const form = document.querySelector('.form');
//Create the list of questions and links to html
const questionList = document.querySelector('#question-list');

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