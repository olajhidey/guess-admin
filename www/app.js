const { createApp } = Vue

const Login = {
    template: `
    <div class="row justify-content-center">
    <div class="col-md-6">
      <div class="card shadow">
        <div class="card-header bg-primary text-white text-center">
          <h4>Login</h4>
        </div>
        <div class="card-body">
          <form @submit.prevent="handleLogin">
            <div v-if="alertMessage" class="alert alert-danger">
                {{ alertMessage }}
            </div>
            <div class="mb-3">
              <label>Username</label>
              <input v-model="username" type="text" class="form-control" required />
            </div>
            <div class="mb-3">
              <label>Password</label>
              <input v-model="password" type="password" class="form-control" required />
            </div>
            <button type="submit" class="btn btn-primary w-100">Login</button>
          </form>
          <div class="text-center mt-3">
            <router-link to="/register">Don't have an account? Register</router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
    `,
    data() {
        return {
            username: '',
            password: '',
            isLoading: false,
            alertMessage: ''
        }
    },
    methods: {
        async handleLogin() {
            try {
                this.isLoading = true;

                const request = await axios.post('/api/auth/login', {
                    username: this.username,
                    password: this.password
                })

                if (request.status === 200) {
                    const response = request.data;
                    this.$router.push('/dashboard');
                    this.saveTokenToLocalStorage(response.token);
                }
            } catch (err) {
                this.isLoading = false;
                console.error(err);
                this.alertMessage = err.response.data.error
            }
        },
        saveTokenToLocalStorage(token) {
            localStorage.setItem('token', token);
        },
    }

}

const Register = {
    template: `
    <div class="row justify-content-center">
    <div class="col-md-6">
      <div class="card shadow">
        <div class="card-header bg-success text-white text-center">
          <h4>Register</h4>
        </div>
        <div class="card-body">
          <div v-if="alertMessage" class="alert alert-danger">
            {{ alertMessage }}
          </div>
          <form @submit.prevent="handleRegister">
            <div class="mb-3">
              <label>Username</label>
              <input v-model="username" type="text" class="form-control" required />
            </div>
            <div class="mb-3">
              <label>Email</label>
              <input v-model="email" type="email" class="form-control" required />
            </div>
            <div class="mb-3">
              <label>Password</label>
              <input v-model="password" type="password" class="form-control" required />
            </div>
            <div class="mb-3">
              <label>Confirm Password</label>
              <input v-model="confirmPassword" type="password" class="form-control" required />
            </div>
            <button type="submit" class="btn btn-success w-100">Register</button>
          </form>
          <div class="text-center mt-3">
            <router-link to="/">Already have an account? Login</router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
    `,
    data() {
        return {
            username: '',
            email: '',
            password: '',
            confirmPassword: '',
            isLoading: false,
            alertMessage: ''
        }
    },

    methods: {
        async handleRegister() {
            try {
                this.isLoading = true;

                const request = await axios.post('/api/auth/register', {
                    username: this.username,
                    email: this.email,
                    password: this.password
                })

                if (request.status === 200) {
                    const response = request.data;
                    this.saveTokenToLocalStorage(response.token);
                    this.$router.push('/login');
                }
            } catch (err) {
                this.isLoading = false;
                console.log(err);
                this.alertMessage = err.response.data.error
            }
        },
        saveTokenToLocalStorage(token) {
            localStorage.setItem('token', token);
        },
    }
}

const DashboardLayout = {
    template: `
    <div>
    <!-- Navbar -->
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark d-lg-none">
      <div class="container-fluid">
        <a class="navbar-brand" href="#">Dashboard</a>
        <button class="navbar-toggler" type="button" data-bs-toggle="offcanvas" data-bs-target="#mobileSidebar">
          <span class="navbar-toggler-icon"></span>
        </button>
      </div>
    </nav>

    <!-- Mobile Offcanvas Sidebar -->
    <div class="offcanvas offcanvas-start text-bg-dark" tabindex="-1" id="mobileSidebar" ref="mobileSidebar">
      <div class="offcanvas-header">
        <h5 class="offcanvas-title">Menu</h5>
        <button type="button" class="btn-close btn-close-white" data-bs-dismiss="offcanvas"></button>
      </div>
      <div class="offcanvas-body">
        <router-link class="nav-link text-white" to="/dashboard" exact @click="closeSidebar">Home</router-link>
        <router-link class="nav-link text-white" to="/dashboard/categories" @click="closeSidebar">Categories</router-link>
        <router-link class="nav-link text-white" to="/dashboard/topics" @click="closeSidebar">Topics</router-link>
        <router-link class="nav-link text-white" to="/dashboard/game/session" @click="closeSidebar">Game Sessions</router-link>
        <!-- <router-link class="nav-link text-white" to="/dashboard/questions" @click="closeSidebar">Questions</router-link> -->
        <a href="#" class="nav-link text-white" @click.prevent="logout" @click="closeSidebar">Logout</a>
      </div>
    </div>

    <!-- Main Layout -->
    <div class="dashboard d-flex">
      <!-- Desktop Sidebar -->
      <div class="sidebar bg-dark text-white d-none d-lg-block p-3" style="width: 250px;">
        <h4 class="text-center">Dashboard</h4>
        <router-link class="nav-link text-white" to="/dashboard" exact>Home</router-link>
        <router-link class="nav-link text-white" to="/dashboard/categories">Categories</router-link>
        <router-link class="nav-link text-white" to="/dashboard/topics">Topics</router-link>
        <router-link class="nav-link text-white" to="/dashboard/game/session">Game Sessions</router-link>
        <a href="#" class="nav-link text-white" @click.prevent="logout">Logout</a>
      </div>

      <!-- Main Content -->
      <div class="content flex-grow-1 p-3 bg-light">
        <router-view></router-view>
      </div>
    </div>
  </div>
  `,
    methods: {
        logout() {
            console.log('Logging out...');
            localStorage.removeItem('token');
            this.$router.push('/login');
        },
        closeSidebar() {
            const sidebarEl = this.$refs.mobileSidebar;
            if (sidebarEl) {
                const offcanvas = bootstrap.Offcanvas.getInstance(sidebarEl);
                if (offcanvas) offcanvas.hide();
            }
        }
    }
}

const DashboardHome = {
    template: `
    <div>
    <h2 class="mb-4">Dashboard Overview</h2>
    <div class="row">
      <div class="col-md-4 mb-3">
        <div class="card text-white bg-primary h-100 shadow">
          <div class="card-body">
            <h5 class="card-title">Total Categories</h5>
            <p class="card-text fs-4">{{totalCategories}}</p>
          </div>
          <div class="card-footer text-end">
            <small>Check logs</small>
          </div>
        </div>
      </div>
      <div class="col-md-4 mb-3">
        <div class="card text-white bg-success h-100 shadow">
          <div class="card-body">
            <h5 class="card-title">Total Questions</h5>
            <p class="card-text fs-4">{{totalQuestions}}</p>
          </div>
          <div class="card-footer text-end">
            <small>Check logs</small>
          </div>
        </div>
      </div>
      <div class="col-md-4 mb-3">
        <div class="card text-white bg-secondary h-100 shadow">
          <div class="card-body">
            <h5 class="card-title">Game Played</h5>
            <p class="card-text fs-4">{{totalGames}}</p>
          </div>
          <div class="card-footer text-end">
            <small>Check logs</small>
          </div>
        </div>
      </div>
      <div class="col-md-4 mb-3">
        <div class="card text-white bg-danger h-100 shadow">
          <div class="card-body">
            <h5 class="card-title">Topics</h5>
            <p class="card-text fs-4">{{ totalTopics }}</p>
          </div>
          <div class="card-footer text-end">
            <small>Check logs</small>
          </div>
        </div>
      </div>

      <div class="col-md-4 mb-3">
          <div class="card text-white bg-warning h-100 shadow">
            <div class="card-body">
              <h5 class="card-title">Users</h5>
              <p class="card-text fs-4">{{totalUsers}}</p>
            </div>
            <div class="card-footer text-end">
              <small>Check Logs</small>
            </div>
          </div>
        </div>
    </div>
  </div>
        `,
    mounted() {
        // console.log('Dashboard Home mounted');
        Promise.all([
            this.fetchCategories(),
            this.fetchQuestions(),
            this.fetchUsers(),
            this.fetchTopics(),
            this.fetchGames()
        ]).catch(err => {
            if(err.status === 401){
                console.log('Unauthorized access. Redirecting to login...');
                window.localStorage.removeItem('token');
                this.$router.push('/login');
            }
            console.error("fetching data err",err);
        })
    },
    data() {
        return {
            categories: [],
            questions: [],
            games: [],
            users: [],
            totalUsers: 0,
            totalCategories: 0,
            totalQuestions: 0,
            totalTopics: 0,
            totalGames: 0
        }
    },
    methods: {
     
        async fetchTopics() {
            try {
                const request = await axios.get('/api/topic/list', {
                    'headers': {
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    }
                });
                this.topics = request.data;
                this.totalTopics = this.topics.length;
            } catch (err) {
              console.log("topic err", err.status);
              if(err.status === 401){
                console.log('Unauthorized access. Redirecting to login...');
                window.localStorage.removeItem('token');
                this.$router.push('/login');
              }
                console.error(err);
            }
        },
        async fetchCategories() {
            try {
                const request = await axios.get('/api/category/list', {
                    'headers': {
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    }
                });
                this.categories = request.data;
                this.totalCategories = this.categories.length;
            } catch (err) {
                console.error(err);
            }
        },
        async fetchQuestions() {
            try {
                const request = await axios.get('/api/question/list', {
                    'headers': {
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    }
                });
                this.questions = request.data;
                this.totalQuestions = this.questions.length;
            } catch (err) {
                console.error(err);
            }
        },
        async fetchUsers() {
            try {
                const request = await axios.get('/api/auth/list');
                this.users = request.data;
                this.totalUsers = this.users.length;
            } catch (err) {
                console.error(err);
            }
        },

        async fetchGames(){
            try{
                const request = await axios.get("/api/game/list", {
                    'headers': {
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    }
                })
                if(request.status === 200){
                    this.games = request.data
                    this.totalGames = this.games.length
                }

            }catch(err){
                console.log(err)
            }
        }
    }

}

const CategoriesPage = {
    template: `
    <div class= "container">
      <h2 class="mb-4">Categories</h2>

      <!-- Add Category Button -->
      <button class="btn btn-primary mb-3" @click="toggleForm">
        {{ showForm ? 'Cancel' : 'Add Category' }}
      </button>

      <!-- Add/Edit Form -->
      <div v-if="showForm" class="card mb-4">
        <div class="card-body">
          <h5 class="card-title">{{ isEditing ? 'Edit Category' : 'Add Category' }}</h5>
          <form @submit.prevent="submitForm">
            <div class="mb-3">
              <label class="form-label">Category Name</label>
              <input type="text" class="form-control" v-model="formData.name" required />
            </div>
            <div class="mb-3">
            <label class="form-label">Category Description</label>
            <input type="text" class="form-control" v-model="formData.description" required />
          </div>
            <button type="submit" class="btn btn-success">
              {{ isEditing ? 'Update' : 'Create' }}
            </button>
          </form>
        </div>
      </div>

      <!-- Categories Table -->
      <div class="table-responsive">
        <table class="table table-bordered align-middle">
          <thead class="table-light">
            <tr>
              <th>Name</th>
              <th>Description</th>
              <th class="text-end">Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-if="categories.length === 0">
              <td colspan="3" class="text-center">No categories found.</td>
            </tr>
            <tr v-for="category in categories" :key="category.id">
              <td>{{ category.name }}</td>
              <td>{{ category.description }}</td>
              <td class="text-end">
                <button class="btn btn-sm btn-warning me-2" @click="editCategory(category)">Edit</button>
                <button class="btn btn-sm btn-danger" @click="deleteCategory(category.ID)">Delete</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    `,
    data() {
        return {
            categories: [],
            showForm: false,
            isEditing: false,
            formData: {
                id: 0,
                name: '',
                description: '',
            }
        }
    },
    mounted() {
        this.fetchCategories();
    },

    methods: {
        async fetchCategories() {
            try {
                const request = await axios.get('/api/category/list', {
                    headers: {
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    }
                });
                this.categories = request.data;
            } catch (err) {
                console.error(err);
            }
        },
        async submitForm() {
            try {
                if (this.isEditing) {
                    const request = await axios.put(`/api/category/update/${this.formData.id}`, this.formData, {
                        headers: {
                            'Authorization': `Bearer ${localStorage.getItem('token')}`
                        }
                    });
                    if (request.status === 200) {
                        this.showForm = false;
                    }
                    this.fetchCategories();

                } else {
                    this.createNewCategoryItem()
                }
                this.showForm = false;
                this.resetForm();
            } catch (err) {
                console.error(err);
            }
        },
        editCategory(category) {
            this.formData.name = category.name
            this.formData.description = category.description;
            this.formData.id = category.ID;
            this.isEditing = true;
            this.showForm = true;
        },
        deleteCategory(id) {
            if (confirm('Are you sure you want to delete this category?')) {
                axios.delete(`/api/category/delete/${id}`, {
                    headers: {
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    }
                })
                    .then(() => {
                        this.fetchCategories();
                    })
                    .catch(err => {
                        console.error(err);
                    });
            }
        },
        async createNewCategoryItem() {
            const request = await axios.post('/api/category/create', this.formData, {
                headers: {
                    "Authorization": `Bearer ${localStorage.getItem('token')}`
                }
            });
            if (request.status === 200) {
                this.showForm = false;
                const response = request.data;
                this.categories.push(response);
            }
        },
        resetForm() {
            this.formData.id = null;
            this.formData.description = '';
            this.formData.name = '';
            this.isEditing = false;
        },
        toggleForm() {
            this.showForm = !this.showForm;
            if (!this.showForm) {
                this.resetForm();
            }
        }
    }
}

const QuestionsPage = {
    template: `
    <div class="container">

    <nav aria-label="breadcrumb" class="mb-3">
    <ol class="breadcrumb">
      <li class="breadcrumb-item">
        <router-link to="/dashboard/topics">Topics</router-link>
      </li>
 
    </ol>
  </nav>

      <h2>Questions</h2>
      <!-- Add Button -->
      <button class="btn btn-primary mb-3" @click="toggleForm">
        {{ showForm ? 'Cancel' : 'Add Question' }}
      </button>

      <!-- Form -->
      <div v-if="showForm" class="card mb-4">
        <div class="card-body">
          <h5 class="card-title">{{ isEditing ? 'Edit Question' : 'Add Question' }}</h5>
          <form @submit.prevent="submitForm">
    
          <!-- Image Source Selection -->
          <div class="row mb-3">
            <div class="col-md-4">
              <label class="form-label">Image Source</label>
              <div class="form-check">
                <input class="form-check-input" type="radio" value="url" v-model="formData.imageType" id="imgUrl">
                <label class="form-check-label" for="imgUrl">Image URL</label>
              </div>
              <div class="form-check">
                <input class="form-check-input" type="radio" value="file" v-model="formData.imageType" id="imgFile">
                <label class="form-check-label" for="imgFile">Upload File</label>
              </div>
            </div>
    
            <div class="col-md-8" v-if="formData.imageType === 'url'">
              <label class="form-label">Image URL</label>
              <input type="url" class="form-control" v-model="formData.image_url">
            </div>
    
            <div class="col-md-8" v-if="formData.imageType === 'file'">
              <label class="form-label">Image File</label>
              <input type="file" class="form-control" @change="handleFileUpload">
            </div>
          </div>
    
          <!-- Options -->
          <div class="row mb-3">
            <div class="col-md-6" v-for="(option, index) in formData.options" :key="index">
              <label class="form-label">Option {{ String.fromCharCode(65 + index) }}</label>
              <input type="text" class="form-control" v-model="formData.options[index]" required>
            </div>
          </div>
    
          <!-- Correct Answer -->
          <div class="mb-3">
            <label class="form-label">Correct Answer</label>
            <select class="form-select" v-model="formData.answer" required>
              <option disabled value="">Select correct option</option>
              <option v-for="(option, index) in formData.options" :value="option">
                Option {{ String.fromCharCode(65 + index) }}: {{ option }}
              </option>
            </select>
          </div>
    
          <button type="submit" class="btn btn-success">
            {{ isEditing ? 'Update' : 'Create' }}
          </button>
        </form>
        </div>
      </div>

      <!-- Table -->
      <div class="table-responsive">
        <table class="table table-bordered align-middle">
          <thead class="table-light">
            <tr>
              <th>Image Url</th>
              <th>Option 1</th>
              <th>Option 2</th>
              <th>Option 3</th>
              <th>Option 4</th>
              <th>Answer</th>
              <th class="text-end">Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-if="questions.length === 0">
              <td colspan="4" class="text-center">No questions available.</td>
            </tr>
            <tr v-for="q in questions" :key="q.ID">
              <td><img :src="q.image_url" class="img-fluid rounded mb-4" style="width:80px; height:50px;"/></td>
              <td>{{ q.option1 }}</td>
              <td>{{ q.option2 }}</td>
              <td>{{ q.option3 }}</td>
              <td>{{ q.option4 }}</td>
              <td>{{ q.answer }}</td>
              <td class="text-end">
                <button class="btn btn-sm btn-warning me-2" @click="editQuestion(q)">Edit</button>
                <button class="btn btn-sm btn-danger" @click="deleteQuestion(q.ID)">Delete</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div >
    `,
    data() {
        return {
            topicId: this.$route.params.id,
            showForm: false,
            isEditing: false,
            categories: [],
            questions: [],
            formData: {
                ID: 0,
                imageType: 'url', // or 'file'
                image_url: '',
                imageFile: null,
                options: ['', '', '', ''],
                answer: '',
            }
        };
    },
    mounted() {
        Promise.all([
            this.fetchQuestions()
        ]).catch(err => {
            console.error(err);
        })
    },
    methods: {
        toggleForm() {
            this.resetForm();
            this.showForm = !this.showForm;
        },
        handleFileUpload(event) {
            this.formData.imageFile = event.target.files[0];
        },
        async submitForm() {
            const question = {
                id: this.isEditing ? this.formData.ID : Date.now(),
                imageType: this.formData.imageType,
                image_url: this.formData.image_url,
                imageFile: this.formData.imageFile,
                option1: this.formData.options[0],
                option2: this.formData.options[1],
                option3: this.formData.options[2],
                option4: this.formData.options[3],
                answer: this.formData.answer,
                topic_id: this.topicId,
            };

            console.log(question);

            if (this.isEditing) {
                this.updateAndEditQuestion(question);
            } else {
               this.addNewQuestion(question);
            }

            this.resetForm();
            this.showForm = false;
        },
        editQuestion(question) {
            this.formData = {
                ...question,
                category: question.category,
                options: [question.option1, question.option2, question.option3, question.option4],
            };
            this.isEditing = true;
            this.showForm = true;
        },
        deleteQuestion(id) {
            try{
                if (confirm('Are you sure you want to delete this question?')) {
                    const request = axios.delete(`/api/question/delete/${id}`, {
                        headers: {
                            'Authorization': `Bearer ${localStorage.getItem('token')}`
                        }
                    });
                    if (request.status === 200) {
                        this.fetchQuestions();
                    }
                }
            }catch (err) {
                console.error(err);
            }
        },
        resetForm() {
            this.formData = {
                id: null,
                text: '',
                imageType: 'url',
                imageUrl: '',
                imageFile: null,
                options: ['', '', '', ''],
                correctAnswer: '',
            };
            this.isEditing = false;
        },
        getCategoryName(id) {
            const cat = this.categories.find(c => c.ID === id);
            return cat ? cat.name : 'Unknown';
        },
        async fetchCategories() {
            try {
                var request = await axios.get('/api/category/list', {
                    'headers': {
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    }
                })
                this.categories = request.data;
            } catch (err) {
                console.error(err);
            }
        },

        async updateAndEditQuestion(question) {
            try{
                console.log(question)
                const request = await axios.put(`/api/question/update/${question.id}`, question, {
                        headers: {
                            'Authorization': `Bearer ${localStorage.getItem('token')}`
                }})

                if (request.status === 200) {
                   this.fetchQuestions();
                }

            }catch (err) {
                console.error(err);
            }
        },

        async addNewQuestion(question) {
            try{
                const request = await axios.post('/api/question/create', question, {
                    headers: {
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    }
                });

                if (request.status === 200) {
                    const response = request.data;
                    this.questions.push(response);
                }

            }catch (err) {
                console.error(err);
            }
        },

        async fetchQuestions() {
            try {
                var request = await axios.get("/api/question/list/"+this.topicId, {
                    'headers': {
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    }
                })
                if(request.status === 200){
                    const response = request.data
                    this.questions = response;
                    console.log(this.questions)
                }
            } catch (error) {
                console.error(err);
            }
        }
    }
}

const TopicsPage = {
    data() {
      return {
        showForm: false,
        isEditing: false,
        topics: [],
        categories: [],
        formData: {
          id: null,
          name: '',
          description: '',
          categoryId: ''
        }
      };
    },
    mounted(){
        Promise.all([
            this.getTopics(),
            this.fetchCategories()
        ]).catch(err=> {
            console.log(err)
        })
    },
    methods: {
      toggleForm() {
        this.resetForm();
        this.showForm = !this.showForm;
      },
      submitForm() {
        const topic = {
          id: this.isEditing ? this.formData.id : Date.now(),
          name: this.formData.name,
          description: this.formData.description,
          category_id: this.formData.categoryId
        };
  
        if (this.isEditing) {
          this.editSelectedTopic(topic)
        } else {
          this.addNewTopic(topic)
        }
  
        this.resetForm();
        this.showForm = false;
      },
      async editTopic(topic) {
        this.formData = { 
            id: topic.ID,
            name: topic.name, 
            description: topic.description, 
            categoryId: topic.category_id
         };
        this.isEditing = true;
        this.showForm = true;
      },
      async editSelectedTopic(topic){
        try{
            const request = await axios.put("/api/topic/update/"+topic.id, topic, {
                'headers': {
                    "Authorization": `Bearer ${window.localStorage.getItem("token")}`
                }
            })

            if (request.status === 200){
                this.getTopics()
            }
        }catch(err){
            console.log(err)
        }
      },
      async addNewTopic(topic){
        try{
            const request = await axios.post("/api/topic/create", topic, {
                'headers': {
                    "Authorization": `Bearer ${window.localStorage.getItem("token")}`
                }
            })
            if (request.status === 200){
                const response = request.data
                this.topics.push(response)
            }
        }catch(err){
            console.log(err)
        }
      },
     async deleteTopic(id) {
        try{
            const request = await axios.delete(`/api/topic/delete/${id}`, {
                'headers': {
                    "Authorization": `Bearer ${window.localStorage.getItem("token")}`
                }
            })

            if(request.status === 200){
                this.getTopics()
            }
        }catch(err){
            console.log(err)
        }
      },
      resetForm() {
        this.formData = {
          id: null,
          name: '',
          description: '',
          categoryId: ''
        };
        this.isEditing = false;
      },
      async getTopics(){
        try{
            const request = await axios.get("/api/topic/list", {
                'headers': {
                    "Authorization": `Bearer ${window.localStorage.getItem("token")}`
                }
            })

            if(request.status === 200){
                const response = request.data
                this.topics = response
            }

        }catch(err){
            console.log(err)
        }
      },

      async fetchCategories() {
        try {
            var request = await axios.get('/api/category/list', {
                'headers': {
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                }
            })
            this.categories = request.data;
        } catch (err) {
            console.error(err);
        }
    },

      getCategoryName(id) {
        const category = this.categories.find(cat => cat.ID === id);
        return category ? category.name : '';
      }
    },
    template: `
      <div>
        <h2 class="mb-4">Topics</h2>
  
        <!-- Add Button -->
        <button class="btn btn-primary mb-3" @click="toggleForm">
          {{ showForm ? 'Cancel' : 'Add Topic' }}
        </button>
  
        <!-- Form -->
        <div v-if="showForm" class="card mb-4">
          <div class="card-body">
            <h5 class="card-title">{{ isEditing ? 'Edit Topic' : 'Add Topic' }}</h5>
            <form @submit.prevent="submitForm">
              <div class="row mb-3">
                <div class="col-md-6">
                  <label class="form-label">Name</label>
                  <input type="text" class="form-control" v-model="formData.name" required />
                </div>
                <div class="col-md-6">
                  <label class="form-label">Category</label>
                  <select class="form-select" v-model="formData.categoryId" required>
                    <option value="" disabled>Select Category</option>
                    <option v-for="cat in categories" :key="cat.ID" :value="cat.ID">
                      {{ cat.name }}
                    </option>
                  </select>
                </div>
              </div>
              <div class="mb-3">
                <label class="form-label">Description</label>
                <textarea class="form-control" v-model="formData.description" rows="2" required></textarea>
              </div>
              <button type="submit" class="btn btn-success">
                {{ isEditing ? 'Update' : 'Create' }}
              </button>
            </form>
          </div>
        </div>
  
        <!-- Topics Table -->
        <div class="table-responsive">
          <table class="table table-bordered align-middle">
            <thead class="table-light">
              <tr>
                <th>Name</th>
                <th>Description</th>
                <th>Category</th>
                <th class="text-end">Actions</th>
              </tr>
            </thead>
            <tbody>
              <tr v-if="topics.length === 0">
                <td colspan="4" class="text-center">No topics available.</td>
              </tr>
              <tr v-if="topics.length > 0" v-for="topic in topics" :key="topic.ID">
                <td>{{ topic.name }}</td>
                <td>{{ topic.description }}</td>
                <td>{{ getCategoryName(topic.category_id) }}</td>
                <td class="text-end">
                  <button class="btn btn-sm btn-warning me-2" @click="editTopic(topic)">Edit</button>
                  <button class="btn btn-sm btn-danger" @click="deleteTopic(topic.ID)">Delete</button>
                  <router-link class="btn btn-sm btn-info" :to="'/dashboard/topics/' + topic.ID + '/questions'">View</router-link>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    `
  };

const GameSessionPage = {
    data() {
      return {
        gameSessions: []
      };
    },
    mounted(){
        Promise.all([
            this.fetchGameSessions()
        ]).catch(err => {
            console.log(err)
        })
    },
    template: `
      <div>
        <h2 class="mb-4">Game Sessions</h2>
  
        <div class="table-responsive">
          <table class="table table-bordered align-middle">
            <thead class="table-light">
              <tr>
                <th>Topic</th>
                <th>Player Name</th>
                <th>Player Score</th>
                <th>Game Code</th>
                <th class="text-end">Actions</th>
              </tr>
            </thead>
            <tbody>
              <tr v-if="gameSessions.length === 0">
                <td colspan="7" class="text-center">No game sessions found.</td>
              </tr>
              <tr v-for="session in gameSessions" :key="session.ID">
                <td>{{ session.topic_id }}</td>
                <td>{{ session.player_name }}</td>
                <td>{{ session.player_score }}</td>
                <td>{{ session.code }}</td>
                <td class="text-end">
                  <!-- Placeholder for future actions -->
                  <button class="btn btn-sm btn-outline-secondary" disabled>View</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    `,
    mounted() {
      this.fetchGameSessions();
    },
    methods: {
        async fetchGameSessions() {
            try{
                const request = await axios.get("/api/game/list", {
                    'headers': {
                        'Authorization': `Bearer ${window.localStorage.getItem('token')}`
                    }
                })
                if(request.status === 200){
                    this.gameSessions = request.data
                }

            }catch(err){
                console.log(err)
            }
        }
    }
  };
  
// Vue Router Setup
const routes = [
    { path: '/login', component: Login },
    { path: '/register', component: Register },
    {
        path: '/dashboard',
        component: DashboardLayout,
        meta: { requiresAuth: true },
        children: [
            { path: '', component: DashboardHome },
            { path: 'categories', component: CategoriesPage },
            { path: 'questions', component: QuestionsPage },
            { path: 'topics', component: TopicsPage },
            { path: 'game/session', component: GameSessionPage },
            { path: 'topics/:id/questions', component: QuestionsPage, props: true },
        ]
    },
    { path: '/:pathMatch(.*)*', redirect: '/dashboard' }
]

const router = VueRouter.createRouter({
    history: VueRouter.createWebHistory(),
    routes
})

router.beforeEach((to, from, next) => {
    const isAuthenticated = localStorage.getItem('token') !== null;

    if (to.meta.requiresAuth && !isAuthenticated) {
        next({ path: '/login' });
    } else {
        next();
    }
})

// Vue app
const app = Vue.createApp({})
app.use(router)
app.mount('#app')


