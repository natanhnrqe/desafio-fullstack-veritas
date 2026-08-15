import { useState, useEffect } from 'react'
import { getTasks, createTask, updateTask, deleteTask } from './services/api'
import Column from './components/Column'
import './App.css'


const COLUMNS = [
  { id: 'todo',        label: 'A Fazer!' },
  { id: 'in_progress', label: 'Em Progresso!' },
  { id: 'done',        label: 'Concluídas!' }
]

export default function App() {

  const [tasks, setTasks]     = useState([])   
  const [loading, setLoading] = useState(true)  
  const [error, setError]     = useState(null)  

  
  useEffect(() => {
    fetchTasks()
  }, []) // 

  const fetchTasks = async () => {
    try {
      setLoading(true)
      setError(null)
      const data = await getTasks()
      setTasks(data)
    } catch (err) {
      setError('Erro ao carregar tarefas')
    } finally {
      setLoading(false)
    }
  }

  
  const handleCreate = async (taskData) => {
    try {
      const newTask = await createTask(taskData)
      setTasks(prev => [newTask, ...prev]) 
    } catch (err) {
      setError('Erro ao criar tarefa')
    }
  }

  
 const handleUpdate = async (id, taskData) => {
  
  setTasks(prev => prev.map(t => t.id === id ? { ...t, ...taskData } : t))
  try {
    const updated = await updateTask(id, taskData)
    setTasks(prev => prev.map(t => t.id === id ? updated : t))
  } catch (err) {
    
    setError('Erro ao atualizar tarefa')
    fetchTasks()
  }
}

 
  const handleDelete = async (id) => {
    try {
      await deleteTask(id)
      setTasks(prev => prev.filter(t => t.id !== id))
    } catch (err) {
      setError('Erro ao deletar tarefa')
    }
  }

 
  const getTasksByStatus = (status) => {
    return tasks.filter(t => t.status === status)
  }

  if (loading) return <div className="loading">Carregando...</div>

  return (
    <div className="app">
      <header className="app-header">
        <h1>Mini Kanban</h1>
      </header>

      {error && (
        <div className="error-banner">
          {error}
          <button onClick={() => setError(null)}>✕</button>
        </div>
      )}

      <main className="board">
        {COLUMNS.map(col => (
          <Column
            key={col.id}
            column={col}
            tasks={getTasksByStatus(col.id)}
            onCreate={handleCreate}
            onUpdate={handleUpdate}
            onDelete={handleDelete}
          />
        ))}
      </main>
    </div>
  )
}