const API_BASE = 'http://localhost:8080/tasks';

document.addEventListener('DOMContentLoaded', () => {
    loadTasks();
    document.getElementById('task-form').addEventListener('submit', addTask);
    document.getElementById('clear-all').addEventListener('click', clearTasks);
});

async function loadTasks() {
    try {
        const response = await fetch(API_BASE);
        if (!response.ok) throw new Error('Failed to load tasks');
        const tasks = await response.json();
        renderTasks(tasks);
    } catch (error) {
        console.error('Error loading tasks:', error);
        alert('Failed to load tasks');
    }
}

function renderTasks(tasks) {
    const taskList = document.getElementById('task-list');
    taskList.innerHTML = '';
    tasks.forEach(task => {
        const li = document.createElement('li');
        li.className = `task-item ${task.IsDone ? 'completed' : ''}`;
        li.innerHTML = `
            <span class="task-text">${task.Text}</span>
            <div class="task-actions">
                <button class="toggle-btn" onclick="toggleTask(${task.ID}, ${!task.IsDone})" title="${task.IsDone ? 'Undo' : 'Complete'}">
                    ${task.IsDone ? '↩' : '✓'}
                </button>
                <button class="edit-btn" onclick="editTask(${task.ID}, '${task.Text}')" title="Edit">✏️</button>
                <button class="delete-btn" onclick="deleteTask(${task.ID})" title="Delete">🗑️</button>
            </div>
        `;
        taskList.appendChild(li);
    });
}

async function addTask(event) {
    event.preventDefault();
    const input = document.getElementById('task-input');
    const text = input.value.trim();
    if (!text) return;

    try {
        const response = await fetch(API_BASE, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ task: text })
        });
        if (!response.ok) throw new Error('Failed to add task');
        input.value = '';
        loadTasks();
    } catch (error) {
        console.error('Error adding task:', error);
        alert('Failed to add task');
    }
}

async function toggleTask(id, isDone) {
    try {
        const response = await fetch(`${API_BASE}/${id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ changes: isDone })
        });
        if (!response.ok) throw new Error('Failed to update task');
        loadTasks();
    } catch (error) {
        console.error('Error updating task:', error);
        alert('Failed to update task');
    }
}

function editTask(id, currentText) {
    const newText = prompt('Edit task:', currentText);
    if (newText && newText.trim() && newText !== currentText) {
        updateTaskText(id, newText.trim());
    }
}

async function updateTaskText(id, text) {
    try {
        const response = await fetch(`${API_BASE}/${id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ changes: text })
        });
        if (!response.ok) throw new Error('Failed to update task');
        loadTasks();
    } catch (error) {
        console.error('Error updating task:', error);
        alert('Failed to update task');
    }
}

async function deleteTask(id) {
    if (!confirm('Are you sure you want to delete this task?')) return;

    try {
        const response = await fetch(`${API_BASE}/${id}`, {
            method: 'DELETE'
        });
        if (!response.ok) throw new Error('Failed to delete task');
        loadTasks();
    } catch (error) {
        console.error('Error deleting task:', error);
        alert('Failed to delete task');
    }
}

async function clearTasks() {
    if (!confirm('Are you sure you want to clear all tasks?')) return;

    try {
        const response = await fetch(API_BASE, {
            method: 'DELETE'
        });
        if (!response.ok) throw new Error('Failed to clear tasks');
        loadTasks();
    } catch (error) {
        console.error('Error clearing tasks:', error);
        alert('Failed to clear tasks');
    }
}