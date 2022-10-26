<script lang="ts">
	export let data: any;
</script>

<div class="h-<svelte:window /> flex items-center justify-center flex-col">
	<h1>Welcome, {data?.profile?.name || ''}!</h1>

	<div class="stats shadow">
		<div class="stat">
		<div class="stat-title">Number of Books</div>
		<div class="stat-value">{data?.bookCount}</div>
		<div class="stat-desc">And more being added!</div>
		</div>
	</div>
</div>
