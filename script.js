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
const books_list= document.querySelector(".books_list");

// Buscar Libro
searchButton.addEventListener("click", async () => {
    try {
        // Realiza la solicitud a la API principal para obtener información del libro
        const response = await axios.get(
            `http://localhost:8080/libros/${searchInput.value}`
        );

        // Cargar Datos Libro
        titleInput.value = response.data["titulo"];
        authorInput.value = response.data["autor"];
        editionInput.value = response.data["edicion"];
        countryInput.value = response.data["pais"];
        dateInput.value = response.data["publicacion"];
        bestSellerYesInput.checked = response.data["bestseller"] === true;
        bestSellerNoInput.checked = response.data["bestseller"] === false;

        window.alert("Se cargaron correctamente los datos del libro");
    } catch (error) {
        
        window.alert("Error al buscar el libro, es posible que ese ID no exista");
    }
});

// Agregar Libros
addBookButton.addEventListener("click", async () => {
    try {
        // Campos de Entrada
        const title = document.getElementById("title_input").value;
        const author = document.getElementById("author_input").value;
        const edition = parseInt(document.getElementById("edition_input").value);
        const country = document.getElementById("country_input").value;
        const date = parseInt(document.getElementById("date_input").value);
        const bestSeller = document.querySelector(".radio input[name='boolean']:checked").value === "on";

        // Información del Libro
        const bookData = {
            titulo: title,
            autor: author,
            edicion: edition,
            pais: country,
            publicacion: date,
            bestseller: bestSeller,
        };

        
        const response = await axios.post(
            `http://localhost:8080/libros`,
            bookData
        );
        listar();

        window.alert("Los datos del libro fueron almacenados correctamente");
    } catch (error) {
        
        window.alert("Error al agregar el libro, verifique los datos e intente nuevamente");
    }
});

// Eliminar Libros
deleteBookButton.addEventListener("click", async () => {
    try {
        
        const response = await axios.delete(
            `http://localhost:8080/libros/${searchInput.value}`
        );

        window.alert("Se eliminó el libro de la base de datos correctamente");
        listar();
    } catch (error) {
        
        window.alert("Error al eliminar el libro, es posible que el ID no sea correcto");
    }
    
});

var cnt=document.querySelector(".books_list");
async function listar (){
searchInput.value="";
try {
	try {
		var ul= cnt.querySelector("ul");
		ul.remove();
	
	}
	catch (error) {
        	console.log("No hay nada");
    }

	var ul=document.createElement("ul");
	var lista=cnt.appendChild(ul);
	
        // Realiza la solicitud a la API principal para obtener información del libro
        
        const response = await axios.get(
            'http://localhost:8080/libros'
        );
        
        

        for (let i=0; i<response['data'].length;i++){
            var list= (response['data'][i]['id'])
            var li=document.createElement("li");
            li.classList.add("libro");
            var span=document.createElement("span");
            span.textContent=list;
            li.appendChild(span);
            lista.appendChild(li);
            li.addEventListener('click', (e)=>{
                const item=e.target;
                searchInput.value=item.querySelector("span").textContent;
                books_list.style.display="none";
            })
            
        }
    } catch (error) {
        
        window.alert("No se agregó la lista");}
};

searchInput.addEventListener('click', (e)=>{
	books_list.style.display="block";
})

listar();

    

// Actualizar
saveBookButton.addEventListener("click", async () => {
    try {
        
        const updatedTitle = titleInput.value;
        const updatedAuthor = authorInput.value;
        const updatedEdition = editionInput.value;
        const updatedCountry = countryInput.value;
        const updatedDate = dateInput.value;
        const updatedBestSeller = bestSellerYesInput.checked;

        //Información actualizada del libro
        const updatedBookData = {
            titulo: updatedTitle,
            autor: updatedAuthor,
            edicion: updatedEdition,
            pais: updatedCountry,
            publicacion: updatedDate,
            bestseller: updatedBestSeller,
        };

        
        const response = await axios.patch(
            `http://localhost:8080/libros/${searchInput.value}`,
            updatedBookData
        );

        window.alert("Se guardaron los cambios correctamente");
    } catch (error) {
       
        window.alert("Error al guardar cambios, verifique los datos e intente nuevamente");
    }
});
