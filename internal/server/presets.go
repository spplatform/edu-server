package server

var presetRoadmaps = map[int][]Milestone{
	1: []Milestone{
		{
			Main:        true,
			Order:       1,
			Description: "Основы программирования",
			CourseID:    0,
			Link:        "https://ru.hexlet.io/courses/programming-basics",
			Steps: []Step{
				{
					Description: "Значение null",
					Link:        "https://ru.hexlet.io/courses/programming-basics/lessons/null/theory_unit",
				},
				{
					Description: "Функции и побочные эффекты",
					Link:        "https://ru.hexlet.io/courses/programming-basics/lessons/pure/theory_unit",
				},
				{
					Description: "Первая программа",
					Link:        "https://ru.hexlet.io/courses/programming-basics/lessons/first_program/theory_unit",
				},
			},
		},
		{
			Main:        true,
			Order:       2,
			Description: "Разработка интерфейсов: вёрстка и JavaScript",
			CourseID:    0,
			Link:        "https://www.coursera.org/specializations/razrabotka-interfeysov",
			Steps: []Step{
				{
					Description: "JavaScript, часть 1: основы и функции",
					Link:        "https://www.coursera.org/learn/javascript-osnovy-i-funktsii?specialization=razrabotka-interfeysov",
				},
				{
					Description: "Основы HTML и CSS",
					Link:        "https://www.coursera.org/learn/snovy-html-i-css?specialization=razrabotka-interfeysov",
				},
				{
					Description: "Тонкости верстки",
					Link:        "https://www.coursera.org/learn/tonkosti-verstki?specialization=razrabotka-interfeysov",
				},
				{
					Description: "JavaScript, часть 2: прототипы и асинхронность",
					Link:        "https://www.coursera.org/learn/javascript-prototipy?specialization=razrabotka-interfeysov",
				},
			},
		},
		{
			Main:        true,
			Order:       3,
			Description: "",
			CourseID:    0,
			Link:        "",
			Steps: []Step{
				{
					Description: "",
					Link:        "",
				},
				{
					Description: "",
					Link:        "",
				},
				{
					Description: "",
					Link:        "",
				},
				{
					Description: "",
					Link:        "",
				},
			},
		},
		{
			Main:        true,
			Order:       4,
			Description: "AWS Fundamentals: Going Cloud-Native",
			CourseID:    0,
			Link:        "https://www.coursera.org/learn/aws-fundamentals-going-cloud-native",
			Steps: []Step{
				{
					Description: "Introduction, Infrastructure, and Compute",
					Link:        "https://www.coursera.org/learn/aws-fundamentals-going-cloud-native",
				},
				{
					Description: "Networking and Storage on AWS",
					Link:        "https://www.coursera.org/learn/aws-fundamentals-going-cloud-native",
				},
				{
					Description: "Databases on AWS",
					Link:        "https://www.coursera.org/learn/aws-fundamentals-going-cloud-native",
				},
				{
					Description: "Monitoring and Scaling",
					Link:        "https://www.coursera.org/learn/aws-fundamentals-going-cloud-native",
				},
			},
		},
	},

	2: []Milestone{
		{
			Main:        true,
			Order:       1,
			Description: "Interaction Design",
			CourseID:    0,
			Link:        "https://www.coursera.org/specializations/interaction-design",
			Steps: []Step{
				{
					Description: "Human-Centered Design: an Introduction",
					Link:        "https://www.coursera.org/learn/human-computer-interaction",
				},
				{
					Description: "Design Principles: an Introduction",
					Link:        "https://www.coursera.org/learn/design-principles",
				},
				{
					Description: "Input and Interaction",
					Link:        "https://www.coursera.org/learn/interaction-techniques",
				},
			},
		},
		{
			Main:        true,
			Order:       2,
			Description: "Растровая графика. Adobe Photoshop CC",
			CourseID:    0,
			Link:        "https://www.coursera.org/learn/rastrovaya-grafika-adobe-photoshop",
			Steps: []Step{
				{
					Description: "Интерфейс. Основы рисования",
					Link:        "https://www.coursera.org/learn/rastrovaya-grafika-adobe-photoshop",
				},
				{
					Description: "",
					Link:        "https://www.coursera.org/learn/rastrovaya-grafika-adobe-photoshop",
				},
				{
					Description: "",
					Link:        "https://www.coursera.org/learn/rastrovaya-grafika-adobe-photoshop",
				},
				{
					Description: "",
					Link:        "https://www.coursera.org/learn/rastrovaya-grafika-adobe-photoshop",
				},
			},
		},
		{
			Main:        true,
			Order:       3,
			Description: "Data Visualization",
			CourseID:    0,
			Link:        "https://www.udacity.com/course/data-visualization-nanodegree--nd197",
			Steps: []Step{
				{
					Description: "Data Visualization",
					Link:        "https://www.udacity.com/course/data-visualization-nanodegree--nd197",
				},
				{
					Description: "Dashboard Designs",
					Link:        "https://www.udacity.com/course/data-visualization-nanodegree--nd197",
				},
				{
					Description: "Data Storytelling",
					Link:        "https://www.udacity.com/course/data-visualization-nanodegree--nd197",
				},
			},
		},
	},
}

var presetInterests = map[int][]Milestone{
	1: []Milestone{
		{
			Main:        false,
			Order:       0,
			Description: "Разработка веб-сервисов на Go - основы языка",
			CourseID:    0,
			Link:        "https://www.coursera.org/learn/golang-webservices-1",
			Steps: []Step{
				{
					Description: "Введение в Golang",
					Link:        "https://www.coursera.org/learn/golang-webservices-1",
				},
				{
					Description: "Асинхронная работа",
					Link:        "https://www.coursera.org/learn/golang-webservices-1",
				},
				{
					Description: "Работа с динамическими данными и производительность",
					Link:        "https://www.coursera.org/learn/golang-webservices-1",
				},
			},
		},
		{
			Main:        false,
			Order:       0,
			Description: "Разработка веб-сервисов на Go - основы языка",
			CourseID:    0,
			Link:        "https://www.coursera.org/learn/golang-webservices-2",
			Steps: []Step{
				{
					Description: "Анатомия веб-сервиса",
					Link:        "https://www.coursera.org/learn/golang-webservices-2",
				},
				{
					Description: "SQL и NoSQL",
					Link:        "https://www.coursera.org/learn/golang-webservices-2",
				},
				{
					Description: "Микросервисы",
					Link:        "https://www.coursera.org/learn/golang-webservices-2",
				},
			},
		},
	},

	2: []Milestone{
		{
			Main:        false,
			Order:       0,
			Description: "Базы данных",
			CourseID:    0,
			Link:        "https://www.coursera.org/learn/data-bases-intr",
			Steps: []Step{
				{
					Description: "Проектирование баз данных. Модель «сущность –связь»",
					Link:        "https://www.coursera.org/learn/data-bases-intr",
				},
				{
					Description: "Реляционная алгебра. Нормализация реляционных отношений",
					Link:        "https://www.coursera.org/learn/data-bases-intr",
				},
				{
					Description: "Основные объекты базы и их описание на языке SQL",
					Link:        "https://www.coursera.org/learn/data-bases-intr",
				},
			},
		},
	},

	3: []Milestone{
		{
			Main:        false,
			Order:       0,
			Description: "Developing Applications with Google Cloud Platform",
			CourseID:    0,
			Link:        "https://www.coursera.org/specializations/developing-apps-gcp",
			Steps: []Step{
				{
					Description: "Google Cloud Platform Fundamentals: Core Infrastructure",
					Link:        "https://www.coursera.org/learn/gcp-fundamentals",
				},
				{
					Description: "Getting Started With Application Development",
					Link:        "https://www.coursera.org/learn/getting-started-app-development",
				},
				{
					Description: "App Deployment, Debugging, and Performance",
					Link:        "https://www.coursera.org/learn/app-deployment-debugging-performance",
				},
			},
		},
	},

	4: []Milestone{
		{
			Main:        false,
			Order:       0,
			Description: "Graphic Design Specialization",
			CourseID:    0,
			Link:        "https://www.coursera.org/specializations/graphic-design",
			Steps: []Step{
				{
					Description: "Fundamentals of Graphic Design",
					Link:        "https://www.coursera.org/learn/fundamentals-of-graphic-design",
				},
				{
					Description: "Introduction to Typography",
					Link:        "https://www.coursera.org/learn/typography",
				},
				{
					Description: "Introduction to Imagemaking",
					Link:        "https://www.coursera.org/learn/image-making",
				},
				{
					Description: "Ideas from the History of Graphic Design",
					Link:        "https://www.coursera.org/learn/graphic-design-history",
				},
			},
		},
	},

	5: []Milestone{
		{
			Main:        false,
			Order:       0,
			Description: "Responsive Web Design",
			CourseID:    0,
			Link:        "https://www.coursera.org/learn/responsive-web-design",
			Steps: []Step{
				{
					Description: "Web design principles",
					Link:        "https://www.coursera.org/learn/responsive-web-design",
				},
				{
					Description: "Realising design principles in code",
					Link:        "https://www.coursera.org/learn/responsive-web-design",
				},
				{
					Description: "Adding content to websites",
					Link:        "https://www.coursera.org/learn/responsive-web-design",
				},
				{
					Description: "Building a full gallery app",
					Link:        "https://www.coursera.org/learn/responsive-web-design",
				},
			},
		},
	},
}
