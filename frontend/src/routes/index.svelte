<script lang="ts">
	import TaskItem from '$lib/TaskItem.svelte';
	import Modal from '$lib/Modal.svelte';
	import { onMount } from 'svelte';
	import { flip } from 'svelte/animate';
	import { push } from '$lib/store';

	// modal
	let showPopup = false;
	let modalTitle = '';
	let modalTaskId = '';

	const onPopupClose = () => {
		showPopup = false;
		modalTaskId = '';
	};
	///
	let inputID = '';
	let inputDetail = '';
	let inputAssignee = '';
	let inputDueDate = '';
	let mode = 'add';

	function inputMode() {
		inputDetail = '';
		inputAssignee = '';
		inputDueDate = '';
		mode = 'add';
	}

	type Task = {
		assignee: string;
		created_at: string;
		deadline: string;
		detail: string;
		done: boolean;
		id: string;
		updated_at: string;
	};

	let tasks: Task[] = [];

	// post new task to api endpoint
	async function addTask() {
		let data = {
			assignee: inputAssignee,
			detail: inputDetail,
			deadline: inputDueDate,
			done: false
		};
		let response = await fetch('/api/tasks', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(data)
		});
		if (response.ok) {
			let json = await response.json();
			let data = json.data;
			push('task created', data.detail, 'success');
			inputMode();
		}

		if (!response.ok) {
			let json = await response.json();
			let data = json.data;
			push('error', data.error);
		}
	}

	// get tasks from api
	async function getTasks() {
		let data = await fetch('/api/tasks')
			.then((r) => r.json())
			.then((r) => r.data);
		tasks = data;
	}

	// set tasks done to api
	async function setTaskDone(id: string, done: boolean) {
		let url: string;
		if (done) {
			url = `/api/tasks/${id}/done`;
		} else {
			url = `/api/tasks/${id}/undone`;
		}
		let data = await fetch(url, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json'
			}
		}).then((r) => r.json());
	}

	// delete task from api
	async function deleteTask(id: string) {
		let url = `/api/tasks/${id}`;
		let data = await fetch(url, {
			method: 'DELETE',
			headers: {
				'Content-Type': 'application/json'
			}
		})
			.then((r) => r.json())
			.then((r) => {
				console.log(r);
			});
	}

	// update task to api
	async function updateTask(id: string) {
		let url = `/api/tasks/${id}`;
		let data = {
			assignee: inputAssignee,
			detail: inputDetail,
			deadline: inputDueDate
		};
		let response = await fetch(url, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(data)
		});
		if (response.ok) {
			let json = await response.json();
			let data = json.data;
			push('task updated', data.detail, 'success');
			inputMode();
		}

		if (!response.ok) {
			let json = await response.json();
			let data = json.data;
			push('error', data.error);
		}
	}

	onMount(async () => {
		await getTasks();
	});

	// handlers
	function handleUpdateTask(task: Task) {
		inputID = task.id;
		inputDetail = task.detail;
		inputAssignee = task.assignee;
		inputDueDate = formatDate(task.deadline);
		mode = 'edit';
	}

	function handleUpdate() {
		updateTask(inputID).then(() => {
			getTasks();
		});
	}

	function handleUpdateCancel() {
		inputMode();
	}

	function handleClearInput() {
		inputMode();
	}

	function handleSetDone(task: Task, status: boolean) {
		setTaskDone(task.id, status).then(() => {
			getTasks();
		});
	}

	function handleDelete(task: Task) {
		showPopup = true;
		modalTaskId = task.id;
		modalTitle = 'Delete Task?';
	}

	function handleDeleteConfimed() {
		deleteTask(modalTaskId).then(() => {
			getTasks();
		});
		showPopup = false;
		modalTaskId = '';
	}

	function handleAddTask() {
		addTask().then(() => {
			getTasks();
		});
	}

	// convert rfc3339 to yyyy-MM-dd
	function formatDate(date: string) {
		let d = new Date(date);
		return d.toISOString().split('T')[0];
	}
</script>

{#if showPopup}
	<Modal
		title={modalTitle}
		open={showPopup}
		onClosed={() => onPopupClose()}
		onOK={handleDeleteConfimed}
	>
		<p>{modalTitle}</p>
	</Modal>
{/if}

<h1>Mini Project Golang DTS</h1>

<div class="row">
	<div class="col">
		<div class="input-group mb-3">
			<input type="text" class="form-control" placeholder="Task Detail" bind:value={inputDetail} />
		</div>
		<div class="input-group mb-3">
			<input type="text" class="form-control" placeholder="Assignee" bind:value={inputAssignee} />
			<span class="input-group-text">@</span>
			<input type="date" class="form-control" placeholder="Date" bind:value={inputDueDate} />
		</div>
		{#if mode === 'add'}
			<button on:click={handleAddTask} class="btn btn-primary">Add</button>
			<button on:click={handleClearInput} class="btn btn-danger">Clear</button>
		{/if}
		{#if mode === 'edit'}
			<button on:click={handleUpdate} class="btn btn-success">Update</button>
			<button on:click={handleUpdateCancel} class="btn btn-danger">Cancel</button>
		{/if}
	</div>
</div>

<div class="row">
	<div class="col">
		<p class="text-center">Double click list to start edit the task</p>
	</div>
</div>

<div class="row mt-4">
	<div class="col">
		<div class="text-center">
			<h2>TODO</h2>
		</div>

		<ul>
			{#each tasks.filter((t) => !t.done) as task (task.id)}
				<div animate:flip>
					<TaskItem
						deadline={task.deadline}
						assignee={task.assignee}
						IsDone={task.done}
						Detail={task.detail}
						onClickDetail={() => handleUpdateTask(task)}
						onClickDone={(status) => handleSetDone(task, status)}
						onClickDelete={() => handleDelete(task)}
					/>
				</div>
			{/each}
		</ul>
	</div>
	<div class="col">
		<div class="text-center">
			<h2>DONE</h2>
		</div>

		<ul>
			{#each tasks.filter((t) => t.done) as task (task.id)}
				<div animate:flip>
					<TaskItem
						deadline={task.deadline}
						assignee={task.assignee}
						IsDone={task.done}
						Detail={task.detail}
						onClickDetail={() => handleUpdateTask(task)}
						onClickDone={(status) => handleSetDone(task, status)}
						onClickDelete={() => handleDelete(task)}
					/>
				</div>
			{/each}
		</ul>
	</div>
</div>
