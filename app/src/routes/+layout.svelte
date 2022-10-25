<script lang="ts">
	import '../lib/app.css';
	export let data;
</script>

<title>Emporium</title>

<div id="body bg-zinc-900">
	<div class="navbar bg-primary rounded" id="apptoolbar">
		<div class="navbar-start">
			<a class="btn btn-ghost text-xl text-white" href="/">The Emporium</a>
		</div>

		<div class="navbar-end">
			<!-- items for navigating the website. youtube temporarily -->
			<a class="btn btn-ghost text-white" href="/">Home</a>
			<a class="btn btn-ghost" href="/library">Library</a>
			<a class="btn btn-ghost" href="https://reddit.com">Settings</a>
			{#if data?.profile}
				<form method="POST" action="/logout">
					<button class="btn">Logout</button>
				</form>
			{:else}
				<a class="btn btn-ghost" href="/login">Login</a>
				<a class="btn btn-ghost" href="/register">Register</a>
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
