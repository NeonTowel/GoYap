<script lang="ts">
	import '../app.css';
</script>

<style>
  :global(body) {
    background-color: #1c1c1c; /* Darker slate gray */
    margin: 0;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
  }

  .layout-container {
    width: 100%;
    height: 100vh;
    background-color: #242424; /* Medium dark gray */
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    font-size: 1rem;
  }
</style>

<div class="layout-container">
  <slot />
</div>