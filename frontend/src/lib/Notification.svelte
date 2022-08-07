<script>
	import { notifications, remove } from './store';
	import { fade } from 'svelte/transition';
	import { flip } from 'svelte/animate';
</script>

<div class="notification-container">
	{#each [...$notifications] as [key, item] (key)}
		<div class="notification" animate:flip in:fade>
			<div class="alert alert-{item.type} alert-dismissible" role="alert">
				<div><strong>{item.title}</strong></div>
				<span>{item.message}</span>
				<button
					on:click={() => {
						remove(key);
					}}
					type="button"
					class="btn-close"
					data-bs-dismiss="alert"
					aria-label="Close"
				/>
			</div>
		</div>
	{/each}
</div>

<style>
	.notification-container {
		display: block;
		position: fixed;
		top: 0;
		right: 0;
		z-index: 10;
	}
	.notification {
		min-width: 400px;
		border-top-left-radius: 0.25rem;
	}

	@media only screen and (max-width: 600px) {
		.notification {
			right: 0;
			left: 0;
			margin: auto;
		}
	}
</style>
