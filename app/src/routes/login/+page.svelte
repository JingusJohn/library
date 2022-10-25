<script lang="ts">
	let email = '';
	let password = '';
</script>

<div class="flex justify-center items-center h-screen">
	<form class="flex flex-col" method="POST" action="?/login">
		<input
			class="input w-full max-w-xs"
			name="email"
			type="text"
			placeholder="email"
			id="email"
			bind:value={email}
		/>
		<input
			class="input w-full max-w-xs"
			name="password"
			type="text"
			placeholder="password"
			id="password"
			bind:value={password}
		/>
		<button class="btn btn-ghost">Login</button>
	</form>
</div>
