import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8080'
})

// Busca
export const getTasks = async () => {
  const response = await api.get('/tasks')
  return response.data
}

// Cria
export const createTask = async (task) => {
  const response = await api.post('/tasks', task)
  return response.data
}

// Atualiza
export const updateTask = async (id, task) => {
  const response = await api.put(`/tasks/${id}`, task)
  return response.data
}

// Deleta
export const deleteTask = async (id) => {
  const response = await api.delete(`/tasks/${id}`)
  return response.data
}