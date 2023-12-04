const searchButton = document.getElementById("search_button");
const addBookButton = document.querySelector("#add_book button");
const deleteBookButton = document.getElementById("delete");
const saveBookButton = document.getElementById("save");
const searchInput = document.getElementById("search_input");
const titleInput = document.getElementById("title_input_search");
const authorInput = document.getElementById("author_input_search");
const editionInput = document.getElementById("edition_input_search");
const countryInput = document.getElementById("country_input_search");
const dateInput = document.getElementById("date_input_search");
const bestSellerYesInput = document.getElementById("best_seller_input_yes_search");
const bestSellerNoInput = document.getElementById("best_seller_input_no_search");

// Event listener para buscar libro
searchButton.addEventListener("click", async () => {
    try {
        // Realiza la solicitud a la API principal para obtener información del libro
        const response = await axios.get(
            `http://localhost:8080/libros/${searchInput.value}`
        );

        // Cargar datos del libro
        titleInput.value = response.data["titulo"];
        authorInput.value = response.data["autor"];
        editionInput.value = response.data["edicion"];
        countryInput.value = response.data["pais"];
        dateInput.value = response.data["fechaPublicacion"];
        bestSellerYesInput.checked = response.data["bestSeller"] === true;
        bestSellerNoInput.checked = response.data["bestSeller"] === false;

        window.alert("Se cargaron correctamente los datos del libro");
    } catch (error) {
        // Manejo de errores
        window.alert("Error al buscar el libro, es posible que ese ID no exista");
    }
});

// Event listener para agregar libro
addBookButton.addEventListener("click", async () => {
    try {
        // Obtener valores de los campos de entrada
        const title = document.getElementById("title_input").value;
        const author = document.getElementById("author_input").value;
        const edition = document.getElementById("edition_input").value;
        const country = document.getElementById("country_input").value;
        const date = document.getElementById("date_input").value;
        const bestSeller = document.querySelector(".radio input[name='boolean']:checked").value === "on";

        // Objeto con la información del libro
        const bookData = {
            titulo: title,
            autor: author,
            edicion: edition,
            pais: country,
            fechaPublicacion: date,
            bestSeller: bestSeller,
        };

        // Realiza la solicitud a la API principal para agregar el libro
        const response = await axios.post(
            `http://localhost:8080/libros`,
            bookData
        );

        window.alert("Los datos del libro fueron almacenados correctamente");
    } catch (error) {
        // Manejo de errores
        window.alert("Error al agregar el libro, verifique los datos e intente nuevamente");
    }
});

// Event listener para eliminar libro
deleteBookButton.addEventListener("click", async () => {
    try {
        // Realiza la solicitud a la API principal para eliminar el libro
        const response = await axios.delete(
            `http://localhost:8080/libros/${searchInput.value}`
        );

        window.alert("Se eliminó el libro de la base de datos correctamente");
    } catch (error) {
        // Manejo de errores
        window.alert("Error al eliminar el libro, es posible que el ID no sea correcto");
    }
});

// Event listener para guardar cambios en la información del libro
saveBookButton.addEventListener("click", async () => {
    try {
        // Obtener valores actualizados de los campos de entrada
        const updatedTitle = titleInput.value;
        const updatedAuthor = authorInput.value;
        const updatedEdition = editionInput.value;
        const updatedCountry = countryInput.value;
        const updatedDate = dateInput.value;
        const updatedBestSeller = bestSellerYesInput.checked;

        // Objeto con la información actualizada del libro
        const updatedBookData = {
            titulo: updatedTitle,
            autor: updatedAuthor,
            edicion: updatedEdition,
            pais: updatedCountry,
            fechaPublicacion: updatedDate,
            bestSeller: updatedBestSeller,
        };

        // Realiza la solicitud a la API principal para actualizar la información del libro
        const response = await axios.put(
            `http://localhost:8080/libros/${searchInput.value}`,
            updatedBookData
        );

        window.alert("Se guardaron los cambios correctamente");
    } catch (error) {
        // Manejo de errores
        window.alert("Error al guardar cambios, verifique los datos e intente nuevamente");
    }
});
