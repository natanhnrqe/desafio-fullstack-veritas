import { useState } from 'react'
import TaskCard from './TaskCard'

export default function Column({ column, tasks, onCreate, onUpdate, onDelete }) {
  const [showForm, setShowForm] = useState(false)
  const [title, setTitle]       = useState('')
  const [description, setDescription] = useState('')
  const [loading, setLoading]   = useState(false)

  const handleSubmit = async (e) => {
    e.preventDefault() 
    if (!title.trim()) return

    setLoading(true)
    await onCreate({
      title,
      description,
      status: column.id 
    })
    // Limpa o formulário
    setTitle('')
    setDescription('')
    setShowForm(false)
    setLoading(false)
  }

  return (
    <div className="column">
      {/* Cabeçalho da coluna */}
      <div className="column-header">
        <h2>{column.label}</h2>
        <span className="task-count">{tasks.length}</span>
      </div>

      {/* Lista de tarefas */}
      <div className="task-list">
        {tasks.map(task => (
          <TaskCard
            key={task.id}
            task={task}
            onUpdate={onUpdate}
            onDelete={onDelete}
          />
        ))}
      </div>

      {/* Formulário de nova tarefa */}
      {showForm ? (
        <form className="task-form" onSubmit={handleSubmit}>
          <input
            type="text"
            placeholder="Título da tarefa"
            value={title}
            onChange={e => setTitle(e.target.value)}
            autoFocus
          />
          <textarea
            placeholder="Descrição (opcional)"
            value={description}
            onChange={e => setDescription(e.target.value)}
            rows={2}
          />
          <div className="form-actions">
            <button type="submit" disabled={loading}>
              {loading ? 'Salvando...' : 'Adicionar'}
            </button>
            <button
              type="button"
              className="btn-cancel"
              onClick={() => setShowForm(false)}
            >
              Cancelar
            </button>
          </div>
        </form>
      ) : (
        <button
          className="btn-add"
          onClick={() => setShowForm(true)}
        >
          + Adicionar tarefa
        </button>
      )}
    </div>
  )
}