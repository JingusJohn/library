<script lang="ts">
  export let data;
</script>

<div class="appToolbar" id="apptoolbar">
  <div class="titleText">
    <a href="/">Jack's Book Emporium</a>
  </div>

  <div id="pageMenu">
    <!-- items for navigating the website. youtube temporarily -->
    <a href="/">Home</a>
    <a href="/library">Library</a>
    <a href="https://reddit.com">Settings</a>
    {#if data?.profile}
    <form method="POST" action="/logout">
      <button>Logout</button>
    </form>
    {:else}
    <a href="/login">Login</a>
    <a href="/register">Register</a>
    {/if}
  </div>
</div>

<slot></slot>

<style lang="scss">
  .appToolbar {
    background-color: #4287f5;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    padding-left: 10px;
    padding-right: 10px;
  }
  
  .titleText {
    color: black;
  }

</style>
