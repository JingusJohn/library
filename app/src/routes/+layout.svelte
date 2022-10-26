<script lang="ts">
	import '../lib/app.css';
	export let data: any;
</script>

<title>Emporium</title>

<div id="body bg-zinc-900">
	<div class="navbar rounded bg-primary" id="apptoolbar">
		<div class="navbar-start">
			<a class="btn-ghost btn text-xl text-white" href="/">The Emporium</a>
		</div>

		<div class="navbar-end">
			<!-- items for navigating the website. youtube temporarily -->
			<a class="btn-ghost btn text-white" href="/">Home</a>
			{#if data?.profile}
				<a class="btn-ghost btn" data-sveltekit-prefetch href="/library">Library</a>
				<a class="btn-ghost btn" href="https://reddit.com">Settings</a>
				<form method="POST" action="/logout">
					<button class="btn">Logout</button>
				</form>
			{:else}
				<a class="btn-ghost btn" href="/login">Login</a>
				<a class="btn-ghost btn" href="/register">Register</a>
			{/if}
		</div>
	</div>

	<slot />
</div>

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
