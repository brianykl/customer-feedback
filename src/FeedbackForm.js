import { useState } from 'react';

function FeedbackForm() {
    const [email, setEmail] = useState('');
    const [category, setCategory] = useState('');
    const [feedbackText, setFeedbackText] = useState('');

    const handleSubmit = async (event) => {
        event.preventDefault();
        const feedbackData = {
            Email: email,         
            Category: category,  
            FeedbackText: feedbackText 
        };

        try {
            // API call to backend
            const response = await fetch('http://localhost:8080/feedback', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(feedbackData)
            });

            if (response.ok) {
                // Handle successful response
                const responseData = await response.json();
                console.log('Response from backend:', responseData);
                // Any additional logic on successful submission
            } else {
                // Handle error
                console.error('Failed to submit feedback:', response.status);
            }
        } catch (error) {
            console.error('Error submitting feedback:', error);
        } finally {
            setEmail('');
            setCategory('');
            setFeedbackText(''); // Reset feedbackText state
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
                    <option value="Service">Service</option>
                    <option value="Product">Product</option>
                    <option value="Pricing">Pricing</option>
                    <option value="Other">Other</option>
                </select>
            </label>
            <label>
                Your feedback:
                <textarea
                    value={feedbackText}
                    onChange={(e) => setFeedbackText(e.target.value)}
                />
            </label>
            <button type="submit">Submit Feedback</button>
        </form>
    );
}

export default FeedbackForm;