<script lang="ts">
  import '../app.css';
</script>

<div class="app">
  <header class="nav-header">
    <a href="/" class="logo">✈️ Triangle Travel</a>
  </header>
  <main class="content">
    <slot />
  </main>
</div>

<style>
  .app {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
  }

  .nav-header {
    display: flex;
    align-items: center;
    padding: 1rem 1.5rem;
    background: var(--surface);
    border-bottom: 1px solid var(--border);
  }

  .logo {
    font-weight: 700;
    font-size: 1.25rem;
    color: var(--text);
    text-decoration: none;
  }

  .logo:hover {
    color: var(--accent);
  }

  .content {
    flex: 1;
  }
</style>
