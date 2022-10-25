<script lang="ts">
  let email = "";
  let password = "";
</script>


<form method="POST" action="?/login">
  <input name="email" type="text" placeholder="email" id="email" bind:value={email}>
  <input name="password" type="text" placeholder="password" id="password" bind:value={password}>
  <button>Login</button>
</form>
