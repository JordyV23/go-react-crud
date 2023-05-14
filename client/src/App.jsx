import "./App.css";

function App() {
  const obtenerDatos = async () => {
    const resp = await fetch("http://localhost:4000/users");
    const data = await resp.json();
    console.log(data);
  };

  return (
    <>
      <div>
        <h1>Hola Mundo</h1>
        <button onClick={obtenerDatos}>Obtener Datos</button>
      </div>
    </>
  );
}

export default App;
