<script>
  import Cookies from "js-cookie";
  let newTask = '';
  let tasks = [];

  function GenerateToken(length) {
    let a = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_".split("");
    let b = [];
    for (let i = 0; i < length; i++) {
      let j = (Math.random() * (a.length - 1)).toFixed(0);
      b[i] = a[j];
    }
    return b.join("");
  }

  async function fetchTasks() {
    let Tasks = [];
    let UserToken = Cookies.get("Token");
    if (UserToken) {
      const response = await fetch(`http://localhost:3000/list/tasks/${UserToken}`);
      if(response.ok){
        const data = await response.json();
        tasks = await data.tasks;
      } else {
        console.log(response)
      }
    } else {
      Cookies.set("Token", GenerateToken(31));
      window.location.href = "/";
    }
  }

  async function createTask() {
    const response = await fetch('http://localhost:3000/create/task', {
      method: 'POST',
      mode: 'cors',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ task: newTask }),
      credentials: 'include',
    });
    if (response.ok) {
      const data = await response.json();
      tasks = [...tasks, data.task];
      newTask = '';
      window.location.href = "/";
    }
  }

  async function deleteTask(token) {
    const response = await fetch(`http://localhost:3000/delete/task/${token}`, {
      method: 'DELETE',
      credentials: 'include',
    });
    if (response.ok) {
      tasks = tasks.filter((task) => task.Token !== token);
    }
  }

  async function updateTaskStatus(token) {
    const response = await fetch(`http://localhost:3000/update/task/${token}`, {
      method: 'PUT',
      credentials: 'include',
    });
    if (response.ok) {
      tasks = tasks.map((task) => {
        if (task.Token === token) {
          task.done = !task.done;
        }
        return task;
      });
    }
  }

  fetchTasks();
</script>

<main>
	<h1>Todo List</h1>

	<form on:submit|preventDefault={createTask}>
    <div class="flex">
		  <input type="text" bind:value={newTask} placeholder="Enter your task..." />
		  <button class="b" type="submit">Add Task</button>
    </div>
	</form>

	<ul>
		{#each tasks as task}
			<li>
				<input type="checkbox" checked={task.Status} on:change={() => updateTaskStatus(task.Token)} />
				<span>{task.Task}</span>
				<button on:click={() => deleteTask(task.Token)}>Delete</button>
			</li>
		{/each}
	</ul>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		font-size: 2em;
	}

	ul {
		padding: 0;
		list-style-type: none;
	}

	li {
		display: flex;
		align-items: center;
		justify-content: space-between;
		margin: 0.5em 0;
	}
  .flex{
    display:flex;
  }
  input[type="text"] {
		border:none;
    margin:0;
    padding:0.5em;
    color:white;
    border-top-left-radius: 8px;
    border-bottom-left-radius: 8px;
	}
	input[type="checkbox"] {
		margin-right: 1em;
	}

	button {
		background-color: #dc3545;
		color: white;
		border: none;
		padding: 0.5em;
		cursor: pointer;
	}
  .b{
    border-radius:0;
    border-top-right-radius: 8px;
    border-bottom-right-radius: 8px;
  }
</style>
