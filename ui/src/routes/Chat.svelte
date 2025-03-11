<script lang="ts">
  import { marked } from 'marked';

  type Message = {
    content: string;
    role: 'user' | 'assistant';
  };

  let messages: Message[] = [];
  let input: string = '';
  let loading: boolean = false;

  function renderMarkdown(content: string) {
    return marked(content);
  }

  async function sendMessage() {
    if (input.trim() === '') return;

    // Add user's message immediately
    messages = [...messages, { content: input, role: 'user' }];

    // Disable input and show loading indicator
    loading = true;
    input = '';

    const contextMessages = messages.slice(-10);
    const response = await fetch('http://localhost:8080/api/chat', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ messages: [...contextMessages, { content: input, role: 'user' }] }),
    });

    try {
      const data = await response.json();
      console.log('Received response:', data);

      if (data.choices && data.choices.length) {
        const assistantMessage = data.choices[0].message;
        messages = [...messages, { content: assistantMessage.content, role: assistantMessage.role }];
      }
    } catch (error) {
      console.error("Failed to parse response:", error);
    }
    
    // Re-enable input
    loading = false;
  }
</script>

<div class="chat-container flex flex-col justify-end bg-dark text-white h-full w-full font-roboto">
  <div class="messages space-y-4 overflow-y-auto flex-grow font-roboto p-6">
    {#each messages as message}
      <div
        class="message p-3 rounded-lg {message.role === 'user' ? 'user-message bg-user text-white' : 'assistant-message text-white'}"
        class:ml-auto={message.role === 'user'}
        >
        {@html message.role === 'assistant' ? renderMarkdown(message.content) : message.content}
      </div>
    {/each}
    {#if loading}
      <!-- Typing indicator -->
      <div class="typing-indicator p-3 text-center">...</div>
    {/if}
  </div>
  <div class="input-container fixed bottom-0 p-4 w-full flex justify-between bg-dark">
    <input class="input bg-input text-white p-4 rounded-lg w-full font-roboto" bind:value={input} on:keydown={e => e.key === 'Enter' && sendMessage()} placeholder={messages.length ? '' : 'How may I help you today?'} disabled={loading} />
  </div>
</div>

<style>
  .chat-container {
    width: 100%;
    height: 100vh;
    position: relative;
  }

  .messages {
    flex-grow: 1;
    overflow-y: auto;
    padding-bottom: 60px;
    padding: 20px;
    font-family: 'Roboto', sans-serif;
  }

  .message {
    font-size: 1rem;
    text-align: left;
  }

  .assistant-message {
    max-width: 80%;
    margin-left: auto;
    margin-right: auto;
    text-align: left;
    padding-left: 20px;
    padding-right: 20px;
  }

  .user-message {
    max-width: 35%;
    align-self: flex-end;
    text-align: left;
  }

  .bg-user {
    background-color: #555;
  }

  .typing-indicator {
    font-size: 1rem;
    color: #999;
  }

  .input-container {
    width: 100%;
    position: sticky;
    bottom: 0;
  }

  .bg-dark { 
    background-color: #1f1f1f; 
  }
  
  .bg-input { 
    background-color: #2b2b2b; 
  }

  /* Optional markdown styles */
  h3 {
    font-size: 1.25rem;
    margin-top: 0.5em;
    margin-bottom: 0.5em;
  }

  ul {
    list-style-type: disc;
    padding-left: 40px;
  }

  li {
    margin-bottom: 0.5em;
  }
</style> 