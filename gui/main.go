package main

import (
	"os"
	"os/exec"
	"regexp"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/yaml.v3"
)

const (
	TEXT     = iota
	PASSWORD = iota
)

type Input struct {
	label        string
	placeholder  string
	inputType    int
	defaultValue string
}

type InputContainer struct {
	container *fyne.Container
	input     *widget.Entry
}

type Values struct {
	Username    string
	Password    string
	Host        string
	Port        int
	ProjectName string
}

var CLIENT_EXE = "arsync_client.exe"

func main() {
	defaultValues := Values{
		Username:    "",
		Password:    "",
		Host:        "",
		Port:        0,
		ProjectName: "",
	}

	// Check if `last_config.yml` exists
	// If it exists, read the values from it
	if _, err := os.Stat("last_config.yml"); err == nil {
		// Read the values from the file
		yamlFile, err := os.ReadFile("last_config.yml")
		if err == nil {
			_ = yaml.Unmarshal(yamlFile, &defaultValues)
		}
	}

	// Create a new app
	app := app.New()
	// Create a new window
	window := app.NewWindow("arsync GUI")
	window.Resize(fyne.NewSize(400, 300))
	window.SetFixedSize(true)

	// Create a quit button
	quitButton := widget.NewButton("Quit", func() {
		app.Quit()
	})

	// If the client exe is not found show an error
	if _, err := os.Stat(CLIENT_EXE); os.IsNotExist(err) {
		errorLabel := widget.NewLabel("Client executable not found. Please put `arsync_client.exe` in the same directory as the GUI.")
		window.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
			errorLabel,
			quitButton,
		))
		window.ShowAndRun()
		return
	}

	// Create inputs and their containers
	inputs := []*InputContainer{
		createInput(&Input{
			label:        "Username",
			placeholder:  "Enter your username",
			inputType:    TEXT,
			defaultValue: defaultValues.Username,
		}),
		createInput(&Input{
			label:        "Password",
			placeholder:  "Enter your password",
			inputType:    PASSWORD,
			defaultValue: defaultValues.Password,
		}),
		createInput(&Input{
			label:        "Host",
			placeholder:  "Enter the host",
			inputType:    TEXT,
			defaultValue: defaultValues.Host,
		}),
		createInput(&Input{
			label:        "Port",
			placeholder:  "Enter the port",
			inputType:    TEXT,
			defaultValue: strconv.Itoa(defaultValues.Port),
		}),
		createInput(&Input{
			label:        "Project name",
			placeholder:  "Enter the project name",
			inputType:    TEXT,
			defaultValue: defaultValues.ProjectName,
		}),
	}

	// Create a download button
	downloadButton := widget.NewButton("Download", func() {
		// Get the values of the inputs
		username := inputs[0].input.Text
		password := inputs[1].input.Text
		host := inputs[2].input.Text
		portStr := inputs[3].input.Text
		projectName := inputs[4].input.Text

		// Validate the inputs
		if username == "" || password == "" || host == "" || portStr == "" || projectName == "" {
			dialog.ShowInformation(
				"Error",
				"All fields are required",
				window,
			)
			return
		}

		// If the port is not a valid number or not in the range of 1-65535 show an error
		port, err := strconv.Atoi(portStr)
		if err != nil || port < 1 || port > 65535 {
			dialog.ShowInformation(
				"Error",
				"Port must be a number between 1 and 65535",
				window,
			)
			return
		}

		// If the project name is not a valid project name show an error
		// Valid project names are Linux folder names without spaces and special characters
		isAlphanumeric := regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(projectName)
		if !isAlphanumeric {
			dialog.ShowInformation(
				"Error",
				"Project name must be alphanumeric",
				window,
			)
			return
		}

		// Update the default values
		defaultValues.Username = username
		defaultValues.Password = password
		defaultValues.Host = host
		defaultValues.Port = port
		defaultValues.ProjectName = projectName

		// Save the values to `last_config.yml`
		d, err := yaml.Marshal(&defaultValues)
		if err != nil {
			dialog.ShowInformation(
				"Error",
				"Could not save last config",
				window,
			)
		}

		err = os.WriteFile("last_config.yml", d, 0644)
		if err != nil {
			dialog.ShowInformation(
				"Error",
				"Could not save last config",
				window,
			)
		}

		// If all inputs are valid, run the download command
		// Run the download command
		exe := "./" + CLIENT_EXE
		command := exec.Command(exe, "-H", host, "-P", portStr, "-u", username, "-p", password, "-f", projectName)
		err = command.Run()
		if err != nil {
			dialog.ShowInformation(
				"Error",
				"An error occurred while downloading the project",
				window,
			)
			return
		}

		dialog.ShowInformation(
			"Success",
			"Project downloaded successfully",
			window,
		)
	})

	// Attach the inputs to the window
	window.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		inputs[0].container,
		inputs[1].container,
		inputs[2].container,
		inputs[3].container,
		inputs[4].container,
		downloadButton,
		quitButton,
	))

	// Run the app
	window.ShowAndRun()
}

func createInput(input *Input) *InputContainer {
	// Create a new input
	var inputWidget *widget.Entry
	switch input.inputType {
	case TEXT:
		inputWidget = widget.NewEntry()
	case PASSWORD:
		inputWidget = widget.NewPasswordEntry()
	}
	// Set the placeholder
	inputWidget.SetPlaceHolder(input.placeholder)
	// Create a new label
	label := widget.NewLabel(input.label)
	// Set the default value
	inputWidget.SetText(input.defaultValue)

	// Create a new container
	container := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		label, inputWidget)

	return &InputContainer{
		container: container,
		input:     inputWidget,
	}
}
