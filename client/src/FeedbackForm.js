import { useState } from 'react';
function FeedbackForm() {
    const[email, setEmail] = useState('');
    const[category, setCategory] = useState('');
    const[feedback, setFeedback] = useState('');

    const handleSubmit = async (event) => {
        event.preventDefault();
        const feedbackData = { 
            email: email, 
            cateogory: category, 
            feedback: feedback};

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
        finally {
            setEmail('');
            setCategory('');
            setFeedback('');    
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <label>
                Email:
                <input
                    type="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
            </label>
            <label>
                Category:
                <select value={category} onChange={(e) => setCategory(e.target.value)}>
                    <option value="">Select a category</option>
                    <option value="service">Service</option>
                    <option value="product">Product</option>
                    <option value="pricing">Pricing</option>
                    <option value="other">Other</option>
                </select>
            </label>
            <label>
                Your feedback:
                <textarea
                    value={feedback}
                    onChange={(e) => setFeedback(e.target.value)}
                />
            </label>
            <button type="submit">Submit Feedback</button>
        </form>
    );
}

export default FeedbackForm;