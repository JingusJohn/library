<script lang="ts">
  export let data;
  export let form;

  let name: string = "";
  let email: string = "";
  let password: string = "";
  let confirmPass: string = "";
  $: validPass = (password == confirmPass && password != "");

</script>


<form method="POST" action="?/register">
  <input name="name" type="text" placeholder="First Name" id="fname" bind:value={name}>
  <input name="email" type="text" placeholder="Email" id="email" bind:value={email}>
  <input name="password" type="text" placeholder="Password" id="password" bind:value={password}>
  <!-- confirm password isn't included in the form submission because it has no name value -->
  <input name="passwordConfirm" type="text" placeholder="Confirm Password" id="confirmPassword" bind:value={confirmPass}>
  <button disabled={!validPass}>Register</button>
</form>
