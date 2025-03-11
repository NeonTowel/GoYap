<script lang="ts">
  import Chat from './Chat.svelte';
</script>

<div class="page-container bg-dark-800 text-white font-roboto h-full w-full flex items-center justify-center">
  <Chat />
</div>

<style>
  .page-container {
    background-color: #282828; /* Dark slate gray */
    font-family: 'Roboto', sans-serif; /* Ensures Roboto will be used */
    font-size: inherit; /* Use inherited font size */
  }

  .h-full {
    height: 100%;
  }

  .w-full {
    width: 100%;
  }
</style>
