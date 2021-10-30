<template>
    <div class="container">
        <h1>GoTodo Web Client</h1>
        <form @submit.prevent="addNewTodo">
            <input v-model="newTodo" type="text" placeholder="New Todo" name="newTodo"/>
            <button>Add</button>
        </form>
        <h2>List</h2>
        <!-- <button @click.prevent="deleteAll">Delete All</button> -->
        <ul>
            <li v-for="todo in todos" :key="todo.id">
                <div v-if="todo.edit">
                    <input v-model="todo.title" class="edit-input" type="text" />
                    <!-- <button @click.prevent="updateTodo(todo.id, todo.title); todo.edit = false;">Update</button> -->
                </div>
                <div v-else>
                    <span :class="{ done: todo.status }" @click="todo.edit = false">
                    {{ todo.title }}
                    </span>
                    <!-- <button @click.prevent="deleteSingle(todo.id)">Delete</button> -->
                </div>
            </li>
        </ul>
    </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Roboto:wght@400;700&display=swap');
* {
    font-family: 'Roboto', sans-serif;
}
body {
    background-color: #222;
}
.container {
    max-width: 800px;
    width: 100%;
    margin: 0 auto;
}
h1, h2 {
    color: #eaeaea;
}
form {
    display: flex;
    margin: 20px 0;
}
input {
    flex: 1;
    font-size: 1.6rem;
    padding: 10px 20px;
    color: #eaeaea;
    background-color: #555;
    border: none;
    outline: none;
}
button {
    font-size: 1.2rem;
    cursor: pointer;
    padding: 10px 20px;
    background-color: #eee;
    border: none;
    outline: none;
}
ul {
    list-style: none;
    padding: 0;
}
li {
    margin: 10px 0;
    padding: 12px;
    color: #eaeaea;
    background-color: #333;
}
li > div {
    display: flex;
    gap: 20px;
    justify-content: space-between;
    align-items: center;
}
.edit-input {
    font-size: 1rem;
    padding: 10px 10px;
}
li > button {
}
li > span.done {
    text-decoration: line-through;
}
</style>

<script>
import axios from 'axios';
export default {
    data () {
        return {
            newTodo: '',
            todos: []
        };
    },
    methods: {
        deleteSingle(id) {
            axios.delete('http://localhost:8000/api/tasks', {
                id: id
            })
            .then(res => {
                this.todos = res.data.data.todos;
                this.todos = this.todos.map(data => ({...data, edit: false}));
            });
        },
        updateTodo(id, title) {
            axios.put('http://localhost:8000/api/tasks', {
                id: id,
                title: title
            })
            .then(res => {
                this.todos = res.data.data.todos;
                this.todos = this.todos.map(data => ({...data, edit: false}));
            });
        },
        addNewTodo() {
            if (this.newTodo.length === 0) return;
            axios.post('http://localhost:8000/api/tasks', {
                title: this.newTodo
            })
            .then(res => {
                this.newTodo = '';
                this.todos = res.data.data.todos;
                this.todos = this.todos.map(data => ({...data, edit: false}));
            });
        }
    },
    created () {
        axios.get('http://localhost:8000/api/tasks')
            .then(res => {
                this.todos = res.data.data.todos;
                // console.log(res.data.data.todos);
                this.todos = this.todos.map(data => ({...data, edit: false}));
            });
    }
}
</script>