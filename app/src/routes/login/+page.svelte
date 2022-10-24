<script lang="ts">
  let email = "";
  let password = "";
</script>


<form>
  <input type="text" placeholder="email" id="email" bind:value={email}>
  <input type="text" placeholder="password" id="password" bind:value={password}>
</form>
