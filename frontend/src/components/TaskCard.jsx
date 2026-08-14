import { useState } from 'react'

const COLUMNS = [
  { id: 'todo',        label: 'A Fazer' },
  { id: 'in_progress', label: 'Em Progresso' },
  { id: 'done',        label: 'Concluídas' }
]

export default function TaskCard({ task, onUpdate, onDelete }) {
  const [isEditing, setIsEditing] = useState(false)
  const [title, setTitle]         = useState(task.title)
  const [description, setDescription] = useState(task.description || '')
  const [loading, setLoading]     = useState(false)

  const handleSave = async () => {
    if (!title.trim()) return
    setLoading(true)
    await onUpdate(task.id, {
      title,
      description,
      status: task.status 
    })
    setIsEditing(false)
    setLoading(false)
  }

  const handleMove = async (newStatus) => {
    setLoading(true)
    await onUpdate(task.id, {
      title: task.title,
      description: task.description || '',
      status: newStatus
    })
    setLoading(false)
  }

  const handleDelete = async () => {
    if (!window.confirm('Deletar essa tarefa?')) return
    setLoading(true)
    await onDelete(task.id)
    setLoading(false)
  }

  
  const currentIndex = COLUMNS.findIndex(c => c.id === task.status)
  const prevColumn   = COLUMNS[currentIndex - 1]
  const nextColumn   = COLUMNS[currentIndex + 1]

  if (isEditing) {
    return (
      <div className="task-card editing">
        <input
          type="text"
          value={title}
          onChange={e => setTitle(e.target.value)}
          autoFocus
        />
        <textarea
          value={description}
          onChange={e => setDescription(e.target.value)}
          placeholder="Descrição (opcional)"
          rows={2}
        />
        <div className="card-actions">
          <button onClick={handleSave} disabled={loading}>
            {loading ? 'Salvando...' : '✓ Salvar'}
          </button>
          <button
            className="btn-cancel"
            onClick={() => {
              setTitle(task.title)
              setDescription(task.description || '')
              setIsEditing(false)
            }}
          >
            ✕ Cancelar
          </button>
        </div>
      </div>
    )
  }

  return (
    <div className={`task-card ${loading ? 'loading' : ''}`}>
      
      <div className="task-content">
        <h3>{task.title}</h3>
        {task.description && (
          <p>{task.description}</p>
        )}
      </div>

      <div className="card-actions">
        {/* Mover para coluna anterior */}
        {prevColumn && (
          <button
            className="btn-move"
            onClick={() => handleMove(prevColumn.id)}
            disabled={loading}
            title={`Mover para ${prevColumn.label}`}
          >
            ← 
          </button>
        )}

      
        <button
        className="btn-edit"
        onClick={() => setIsEditing(true)}
        disabled={loading}
        title="Editar"
        >
        <img src="/edit.svg" alt="Editar" />
        </button>

        
        <button
        className="btn-delete"
        onClick={handleDelete}
        disabled={loading}
        title="Deletar"
        >
        <img src="/trash.svg" alt="Deletar" />
        </button>

        {nextColumn && (
          <button
            className="btn-move"
            onClick={() => handleMove(nextColumn.id)}
            disabled={loading}
            title={`Mover para ${nextColumn.label}`}
          >
            →
          </button>
        )}
      </div>
    </div>
  )
}