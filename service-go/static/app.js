// app.js
document.addEventListener('DOMContentLoaded', () => {
    fetchCats();

    document.getElementById('catForm').addEventListener('submit', handleSubmit);
});

const fetchCats = async () => {
    try {
        const response = await fetch(`${CONFIG.API_BASE_URL}/cats`);
        const cats = await response.json();
        renderCatList(cats);
    } catch (error) {
        alert('Error fetching cats: ' + error);
    }
};

const renderCatList = (cats) => {
    const catList = document.getElementById('catList');
    catList.innerHTML = cats.map(cat => createCatCard(cat)).join('');
};

const createCatCard = (cat) => `
    <div class="cat-card">
        <h3>${cat.name} ${cat.nickName ? `(${cat.nickName})` : ''}</h3>
        <div class="cat-info">性别: ${cat.sex === 'male' ? '公猫' : '母猫'}</div>
        <div class="cat-info">状态: ${getStatusText(cat.status)}</div>
        <div class="cat-info">地理范围: ${cat.area || '未知'}</div>
        <div class="cat-info">具体位置: ${cat.location || '未知'}</div>
        <div style="margin-top: 10px">
            <button class="btn" onclick="editCat('${cat.id}')">编辑</button>
            <button class="btn" onclick="deleteCat('${cat.id}')" style="background-color: #f44336;">删除</button>
        </div>
    </div>
`;

const getStatusText = (status) => ({
    trap: '已捕获',
    neuter: '已绝育',
    return: '已放归'
}[status] || status);

const showAddCatModal = () => {
    document.getElementById('modalTitle').textContent = '添加猫咪';
    document.getElementById('catForm').reset();
    document.getElementById('catModal').style.display = 'block';
};

const closeModal = () => {
    document.getElementById('catModal').style.display = 'none';
};

const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = getFormData();
    try {
        await fetch(`${CONFIG.API_BASE_URL}/cats`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(formData)
        });
        closeModal();
        await fetchCats();
    } catch (error) {
        alert('Error saving cat: ' + error);
    }
};

const getFormData = () => ({
    name: document.getElementById('name').value,
    nickName: document.getElementById('nickName').value,
    sex: document.getElementById('sex').value,
    status: document.getElementById('status').value,
    area: document.getElementById('area').value,
    location: document.getElementById('location').value
});

const editCat = (catId) => {
    const cat = currentCats.find(c => c.id === catId);
    if (cat) {
        populateModal(cat);
        document.getElementById('catModal').style.display = 'block';
    }
};

const deleteCat = async (catId) => {
    if (confirm('确定要删除这只猫咪吗？')) {
        try {
            await fetch(`${CONFIG.API_BASE_URL}/cats/${catId}`, { method: 'DELETE' });
            await fetchCats();
        } catch (error) {
            alert('Error deleting cat: ' + error);
        }
    }
};
