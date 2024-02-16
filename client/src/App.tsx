import React, {useEffect, useState} from 'react';

const serverUrl = "http://localhost:5000"

interface Todo {
  id?: number
  label: string
  checked: boolean
}

function App() {

  const [todos, setTodos] = useState<Todo[]>([])
  const [activeTodo, setActiveTodo] = useState<Todo>({label: "", checked: false})

  const styles = {
    todoWrapper: {
      maxWidth: "100%",
      display: "flex",
      justifyContent: "space-between",
      padding: 24,
      border: "1px solid black"
    },
    actionWrapper: {
      display: "flex",
      gap: "24px"
    }
  }

  // Load todos on render
  useEffect(() => {
    loadTodos().then()
  }, [])

  const loadTodos = async () => {

    let _todos
    let ok = false

    try {
      const res = await fetch(serverUrl + "/todos/", {
        method: "GET",
        mode: "cors",
        headers: {"Content-Type": "application/json"},
      })

      _todos = await res.json()
      ok = res.ok
    } catch (e) {
      console.log(e)
    }

    if (ok) setTodos(_todos)

  }

  const checkTodo = async (todo: Todo) => {

    let _todo = {
      ...todo,
      checked: !todo.checked
    }

    let ok = false

    try {
      const res = await fetch(serverUrl + "/todos/", {
        method: "PUT",
        mode: "cors",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(_todo)
      })

      ok = res.ok
    } catch (e) {
      console.log(e)
    }

    if (ok) {
      loadTodos().then()
    }
  }

  const deleteTodo = async (id: any) => {
    let ok = false

    if (!id) return

    try {
      const res = await fetch(serverUrl + "/todos/" + id, {
        method: "DELETE",
        mode: "cors",
        headers: {"Content-Type": "application/json"},
      })
      ok = res.ok
    } catch (e) {console.log(e)}

    if (ok) loadTodos().then()
  }

  const createTodo = async () => {
    let ok = false

    try {
      const res = await fetch(serverUrl + "/todos/", {
        method: "POST",
        mode: "cors",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(activeTodo)
      })

      ok = res.status === 201
    } catch (e) {
      console.log(e)
    }

    if (ok) loadTodos().then()
  }

  // RENDERS
  const renderTodoEditor = () => {

    const handleChangeTodoName = (e: any) => {
      setActiveTodo({...activeTodo, label: e.target.value})
    }

    return (
        <div style={{margin: "24px 0", ...styles.todoWrapper}}>
          <div>
            <input value={activeTodo.label} onChange={handleChangeTodoName} />
          </div>
          <div>
            <button onClick={() => createTodo()}>
              Confirm
            </button>
          </div>
        </div>
    )
  }

  return (
      <div style={{maxWidth: "100%", display: "block"}}>
        <h1>Todos app</h1>

        {renderTodoEditor()}

        {
          !!todos && todos.map((todo: Todo, index: number) => (
              <div key={index} style={styles.todoWrapper}>
                <div>{todo.label}</div>
                <div style={styles.actionWrapper}>
                  <div>
                    <button onClick={() => checkTodo(todo)}>
                      {todo.checked ? "Unckeck" : "Check"}
                    </button>
                  </div>
                  <div>
                    <button onClick={() => deleteTodo(todo.id)}>Delete</button>
                  </div>
                </div>
              </div>
            ))
        }
    </div>
  );
}

export default App;
