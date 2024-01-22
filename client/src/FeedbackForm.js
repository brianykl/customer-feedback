import { useState } from 'react';
function FeedbackForm() {
    const[feedback, setFeedback] = useState('');

    const handleSubmit = async (event) => {
        event.preventDefault();
        // Add logic to handle the submission, like sending it to a backend
        const feedbackData = { feedback: feedback};
        try {
        // api call to backend
            const response = await fetch('http://127.0.0.1:5000/api/feedback', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(feedbackData)
            });

            if(response.ok) {
                // handle successful response
                const responseData = await response.json();
                console.log('response from backend:', responseData);
                // any additional logic on successful submission
            }

            else {
                // handle error
                console.error('failed to submit feedback:', response.status);
            }
        }
        catch (error) {
            console.error('error submitting feedback:', error);
        }
    };

    return (
        <form onSubmit = {handleSubmit}>
            <label>
                Your feedback:
                <textarea value = {feedback} onChange = {(e) => setFeedback(e.target.value)} />
            </label>
            <button type = "submit" > Submit Feedback</button>
        </form>
    );
}

export default FeedbackForm;