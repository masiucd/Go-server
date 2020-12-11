fetch("http://localhost:8080/")
  .then(r => r.json())
  .then(r => {
    console.log(r)
  })
