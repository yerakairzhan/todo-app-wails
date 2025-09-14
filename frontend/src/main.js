// frontend/src/main.js
import './style.css';

// Global variables
let tasks = [];
let currentFilter = 'all';
let taskToDelete = null;

// Theme management
const themeToggle = document.getElementById('theme-toggle');
const savedTheme = localStorage.getItem('theme') || 'light';
document.documentElement.setAttribute('data-theme', savedTheme);
updateThemeIcon(savedTheme);

// DOM elements
const taskInput = document.getElementById('task-input');
const addTaskBtn = document.getElementById('add-task-btn');
const inputError = document.getElementById('input-error');
const filterTabs = document.querySelectorAll('.tab-btn');
const activeTasksSection = document.getElementById('active-tasks-section');
const completedTasksSection = document.getElementById('completed-tasks-section');
const activeTasks = document.getElementById('active-tasks');
const completedTasks = document.getElementById('completed-tasks');
const emptyState = document.getElementById('empty-state');
const deleteModal = document.getElementById('delete-modal');
const confirmDeleteBtn = document.getElementById('confirm-delete');
const cancelDeleteBtn = document.getElementById('cancel-delete');

// Event listeners
themeToggle.addEventListener('click', toggleTheme);
addTaskBtn.addEventListener('click', addTask);
taskInput.addEventListener('keypress', (e) => {
    if (e.key === 'Enter') addTask();
});
taskInput.addEventListener('input', clearError);

filterTabs.forEach(tab => {
    tab.addEventListener('click', () => setFilter(tab.dataset.filter));
});

confirmDeleteBtn.addEventListener('click', confirmDelete);
cancelDeleteBtn.addEventListener('click', cancelDelete);
deleteModal.addEventListener('click', (e) => {
    if (e.target === deleteModal) cancelDelete();
});

// Initialize app
document.addEventListener('DOMContentLoaded', () => {
    // Дожидаемся загрузки Wails
    if (window.go) {
        loadTasks();
    } else {
        window.addEventListener('DOMContentLoaded', loadTasks);
        // Если window.go недоступен, используем fallback
        setTimeout(loadTasks, 1000);
    }
});

// Theme functions
function toggleTheme() {
    const currentTheme = document.documentElement.getAttribute('data-theme');
    const newTheme = currentTheme === 'dark' ? 'light' : 'dark';

    document.documentElement.setAttribute('data-theme', newTheme);
    localStorage.setItem('theme', newTheme);
    updateThemeIcon(newTheme);
}

function updateThemeIcon(theme) {
    themeToggle.textContent = theme === 'dark' ? '☀️' : '🌙';
}

// Backend API calls using Wails bindings
async function callBackend(method, ...args) {
    try {
        console.log(`Calling backend method: ${method} with args:`, args);

        switch(method) {
            case 'AddTask':
                return await AddTask(...args);
            case 'GetAllTasks':
                return await GetAllTasks(...args);
            case 'ToggleTask':
                return await ToggleTask(...args);
            case 'DeleteTask':
                return await DeleteTask(...args);
            default:
                console.warn(`Unknown backend method: ${method}`);
                return null;
        }
    } catch (error) {
        console.error(`Backend call failed for ${method}:`, error);
        throw error;
    }
}

// Task functions
async function loadTasks() {
    try {
        const result = await callBackend('GetAllTasks');
        if (result) {
            tasks = result;
        } else {
            // Fallback to mock data if backend not available
            tasks = [
                { id: 1, title: "Sample task 1", is_completed: false, created_at: new Date().toISOString() },
                { id: 2, title: "Sample completed task", is_completed: true, created_at: new Date().toISOString() }
            ];
        }
        renderTasks();
    } catch (error) {
        console.error('Failed to load tasks:', error);
        showError('Failed to load tasks');
        // Use mock data as fallback
        tasks = [];
        renderTasks();
    }
}

async function addTask() {
    const title = taskInput.value.trim();

    if (!title) {
        showError('Please enter a task title');
        return;
    }

    if (title.length > 255) {
        showError('Task title is too long (max 255 characters)');
        return;
    }

    try {
        console.log('Adding task:', title);
        const result = await callBackend('AddTask', title);
        console.log('Add task result:', result);

        if (result) {
            // Reload tasks after successful addition
            await loadTasks();
            taskInput.value = '';
            clearError();

            // Show success message
            taskInput.placeholder = 'Task added successfully!';
            setTimeout(() => {
                taskInput.placeholder = 'Add a new task...';
            }, 2000);
        } else {
            // Fallback to local behavior if backend not available
            console.log('Backend not available, adding task locally');
            const newTask = {
                id: Date.now(),
                title: title,
                is_completed: false,
                created_at: new Date().toISOString()
            };
            tasks.unshift(newTask);
            taskInput.value = '';
            clearError();
            renderTasks();

            taskInput.placeholder = 'Task added locally!';
            setTimeout(() => {
                taskInput.placeholder = 'Add a new task...';
            }, 2000);
        }
    } catch (error) {
        console.error('Failed to add task:', error);
        showError('Failed to add task: ' + (error.message || 'Unknown error'));
    }
}

async function toggleTask(id) {
    try {
        const result = await callBackend('ToggleTask', parseInt(id));
        if (result) {
            // Reload tasks after successful toggle
            await loadTasks();
        } else {
            // Fallback to mock behavior
            const task = tasks.find(t => t.id == id);
            if (task) {
                task.is_completed = !task.is_completed;
                renderTasks();
            }
        }
    } catch (error) {
        console.error('Failed to toggle task:', error);
        showError('Failed to update task');
    }
}

function showDeleteModal(id) {
    taskToDelete = parseInt(id);
    deleteModal.classList.add('show');
}

function cancelDelete() {
    taskToDelete = null;
    deleteModal.classList.remove('show');
}

async function confirmDelete() {
    if (!taskToDelete) return;

    try {
        const result = await callBackend('DeleteTask', taskToDelete);

        // Check if backend call was successful (no error thrown)
        if (window.go && window.go.main && window.go.main.App) {
            // Reload tasks after successful deletion
            await loadTasks();
        } else {
            // Fallback to mock behavior
            tasks = tasks.filter(t => t.id != taskToDelete);
            renderTasks();
        }

        cancelDelete();
    } catch (error) {
        console.error('Failed to delete task:', error);
        showError('Failed to delete task');
    }
}

function setFilter(filter) {
    currentFilter = filter;

    // Update active tab
    filterTabs.forEach(tab => {
        tab.classList.toggle('active', tab.dataset.filter === filter);
    });

    renderTasks();
}

function renderTasks() {
    const activeTasks = tasks.filter(task => !task.is_completed);
    const completedTasks = tasks.filter(task => task.is_completed);

    // Clear containers
    document.getElementById('active-tasks').innerHTML = '';
    document.getElementById('completed-tasks').innerHTML = '';

    // Show/hide sections based on filter
    if (currentFilter === 'all') {
        activeTasksSection.classList.remove('hidden');
        completedTasksSection.classList.remove('hidden');
        renderTaskList(activeTasks, 'active-tasks');
        renderTaskList(completedTasks, 'completed-tasks');
    } else if (currentFilter === 'active') {
        activeTasksSection.classList.remove('hidden');
        completedTasksSection.classList.add('hidden');
        renderTaskList(activeTasks, 'active-tasks');
    } else if (currentFilter === 'completed') {
        activeTasksSection.classList.add('hidden');
        completedTasksSection.classList.remove('hidden');
        renderTaskList(completedTasks, 'completed-tasks');
    }

    // Show empty state if no tasks
    const hasVisibleTasks = (currentFilter === 'all' && tasks.length > 0) ||
        (currentFilter === 'active' && activeTasks.length > 0) ||
        (currentFilter === 'completed' && completedTasks.length > 0);

    emptyState.style.display = hasVisibleTasks ? 'none' : 'block';
}

function renderTaskList(taskList, containerId) {
    const container = document.getElementById(containerId);

    taskList.forEach(task => {
        const taskElement = createTaskElement(task);
        container.appendChild(taskElement);
    });
}

function createTaskElement(task) {
    const taskDiv = document.createElement('div');
    taskDiv.className = `task-item ${task.is_completed ? 'completed' : ''}`;

    taskDiv.innerHTML = `
        <input 
            type="checkbox" 
            class="task-checkbox" 
            ${task.is_completed ? 'checked' : ''}
            onchange="toggleTask(${task.id})"
        />
        <span class="task-title">${escapeHtml(task.title)}</span>
        <div class="task-actions">
            <button class="task-btn delete" onclick="showDeleteModal(${task.id})" title="Delete task">
                🗑️
            </button>
        </div>
    `;

    return taskDiv;
}

// Utility functions
function showError(message) {
    inputError.textContent = message;
    inputError.classList.add('show');
}

function clearError() {
    inputError.classList.remove('show');
}

function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}

// Make functions global for onclick handlers
window.toggleTask = toggleTask;
window.showDeleteModal = showDeleteModal;