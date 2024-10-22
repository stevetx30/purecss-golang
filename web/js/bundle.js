document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('contact-form');
    const messageBanner = document.getElementById('message-banner');
  
    // Function to show the message banner
    function showMessage(type, message) {
      messageBanner.textContent = message;
      if (type === 'success') {
        messageBanner.className = 'message-banner message-success';
      } else {
        messageBanner.className = 'message-banner message-error';
      }
      messageBanner.style.display = 'block';
    }
  
    // Form submit event listener
    form.addEventListener('submit', async (event) => {
      event.preventDefault();  // Prevent the form from submitting the default way
  
      // Get form values
      const name = document.getElementById('name').value.trim();
      const email = document.getElementById('email').value.trim();
      const message = document.getElementById('message').value.trim();
  
      // Basic validation
      if (!name || !email || !message) {
        showMessage('error', 'Please fill in all fields.');
        return;
      }
  
      // Prepare data to send
      const formData = {
        name: name,
        email: email,
        message: message,
      };
  
      try {
        // Send data to the API
        const response = await fetch('/api/contact', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(formData),
        });
  
        // Handle the response
        if (response.ok) {
          showMessage('success', 'Form submitted successfully!');
          form.reset();  // Clear the form
        } else {
          showMessage('error', 'Failed to submit the form. Please try again.');
        }
      } catch (error) {
        console.error('Error submitting form:', error);
        showMessage('error', 'An error occurred. Please try again.');
      }
    });
  });
  