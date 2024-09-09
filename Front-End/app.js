const API_URL = 'http://localhost:8080/personagens';

// Função para adicionar um personagem quando o formulário for submetido
document.getElementById('addCharacterForm').addEventListener('submit', async (e) => {
  e.preventDefault();

  // Coleta os dados do formulário
  const data = {
    name: document.getElementById('name').value,
    status: document.getElementById('status').value,
    species: document.getElementById('species').value,
    type: document.getElementById('type').value || null,
    gender: document.getElementById('gender').value,
    image: document.getElementById('image').value,
    url: document.getElementById('url').value
  };

  try {
    // Envia os dados para a API usando o método POST
    const response = await fetch(API_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json' // Define o tipo de conteúdo como JSON
      },
      body: JSON.stringify(data) // Converte os dados para o formato JSON
    });

    const result = await response.json();
    alert('Personagem Adicionado: ' + JSON.stringify(result));
    listCharacter();
  } catch (error) {
    console.error('Erro ao adicionar personagem:', error);
  }
});

// Função para listar os personagens com opção de ordenação ou filtro
async function listCharacter(order = '') {
  let url = API_URL;

  // Filtro de status ou ordenação, adicionando parâmetros à URL
  if (order === 'Alive' || order === 'Dead') {
    url += `?status=${order}`; // Filtro por status
  } else if (order === 'asc' || order === 'desc') {
    url += `?order=${order}`; // Ordenação por ordem ascendente ou descendente
  }

  try {
    // Busca os personagens da API
    const response = await fetch(url);
    const characters = await response.json();

    const list = document.getElementById('characterList');
    list.innerHTML = '';
    characters.forEach(character => {
      const li = document.createElement('li');
      li.textContent = `${character.name} - ${character.species} (${character.status})`;
      list.appendChild(li); // Adiciona cada personagem na lista
    });
  } catch (error) {
    console.error('Erro ao listar personagens', error);
  }
}

// Função para buscar um personagem específico pelo ID
async function searchCharacter() {
  const id = document.getElementById('searchId').value; // Obtém o ID do campo de busca
  const url = `${API_URL}/${id}`; // Monta a URL com o ID do personagem

  try {
    // Busca o personagem da API pelo ID
    const response = await fetch(url);
    const character = await response.json();

    // Exibe os detalhes do personagem encontrado
    const detail = document.getElementById('characterDetail');
    detail.innerHTML = character ? `
            <h3>${character.name}</h3>
            <p>Status: ${character.status}</p>
            <p>Species: ${character.species}</p>
            <p>Type: ${character.type || 'N/A'}</p>
            <p>Gender: ${character.gender}</p>
            <p>Image: <img src="${character.image}" alt="${character.name}" width="100"></p>
            <p>URL: <a href="${character.url}" target="_blank">${character.url}</a></p>
        ` : 'Character not found';
  } catch (error) {
    console.error('Error fetching character:', error);
  }
}

// Função para deletar um personagem pelo ID
async function deleteCharacter() {
  const id = document.getElementById('deleteId').value; // Obtém o ID do campo de deletar
  const url = `${API_URL}/${id}`; // Monta a URL com o ID do personagem

  try {
    // Faz a requisição DELETE para a API
    const response = await fetch(url, {
      method: 'DELETE'
    });

    if (response.ok) {
      alert(`Personagem com ID ${id} foi deletado com sucesso.`);
      listCharacter(); // Atualiza a lista de personagens após a exclusão
    } else {
      const errorData = await response.json();
      alert(`Erro ao deletar personagem: ${errorData.message}`);
    }
  } catch (error) {
    console.error('Erro ao deletar personagem:', error);
  }
}

