# Django Todo App

This is a simple todo app built with Django.

## Installation

1. Clone the repository:
```
git clone https://github.com/aenkya/playground/pie/pecan
```


2. Create a virtual environment and activate it:
```
python -m venv env source env/bin/activate
```

3. Install the dependencies:
```
pip install -r requirements.txt
```

4. Set up the database:
```
python manage.py migrate
```

5. (Optional) Load some initial data:
```
python manage.py loaddata initial_data.json
```

6. Run the development server:
```
python manage.py runserver
```

The app should now be available at http://localhost:8000/.

## Usage

To use the app, simply navigate to http://localhost:8000/ in your web browser. You should see a list of your todos, along with options to add, edit, and delete todos.

## Contributing

If you would like to contribute to the project, you can follow these steps:

1. Fork the repository.

2. Create a new branch for your feature or bug fix:

```
git checkout -b my-feature-branch
```

3. Make your changes and commit them:
```
git commit -m "Add new feature"
```

4. Push your changes to your fork:
```
git push origin my-feature-branch
```

5. Create a pull request on the original repository.

## License

This project is licensed under the MIT License. See the LICENSE file for details.