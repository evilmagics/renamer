# 📁 GO Renamer 🛠️

This Go project provides a command-line tool named `renamer` 🚀 for renaming multiple files within a specified directory by adding either a prefix or a suffix to their existing names. 📝

## ✨ Features ✨

* **Bulk Renaming:** 💨 Efficiently renames a large number of files.
* **Prefix Addition:** ➕ Adds a specified prefix to the beginning of filenames.
* **Suffix Addition:** ➡️ Adds a specified suffix to the end of filenames, correctly handling file extensions. 🧩
* **Command-Line Interface:** ⌨️ Easy to use with command-line arguments.
* **Directory Selection:** 📂 Allows users to specify the target directory for renaming.

## 🚀 Usage 🚀

To use the program, you need to compile the Go source code. Once compiled, you can run the executable with the following command-line arguments:

```bash
./renamer -dir <directory> [-prefix <prefix>] [-suffix <suffix>]
```

- -dir <directory>: 📂 Specifies the directory containing the files to be renamed. This argument is required. ⚠️
- -prefix <prefix>: ➕ Specifies the prefix to be added to the filenames. This argument is optional. ➡️
- -suffix <suffix>: ➡️ Specifies the suffix to be added to the filenames. This argument is optional. ➕

## 🚀 Examples 🚀

Here are some examples of how to use the `renamer` command-line tool:

1.  **Adding a prefix to all files in a directory:**

    ```bash
    ./renamer -dir /path/to/my/files -prefix "new_"
    ```

    This command will rename all files in the `/path/to/my/files` directory by adding the prefix "new\_" to the beginning of each filename. For example, `file.txt` would become `new_file.txt`.

2.  **Adding a suffix to all files in a directory:**

    ```bash
    ./renamer -dir /path/to/my/files -suffix "_renamed"
    ```

    This command will rename all files in the `/path/to/my/files` directory by adding the suffix "\_renamed" to the end of each filename, before the file extension. For example, `image.jpg` would become `image_renamed.jpg`.

3.  **Adding both a prefix and a suffix:**

    ```bash
    ./renamer -dir /path/to/my/files -prefix "processed_" -suffix "_done"
    ```

    This command will rename all files in the `/path/to/my/files` directory by adding the prefix "processed\_" to the beginning and the suffix "\_done" to the end of each filename, before the file extension. For example, `document.pdf` would become `processed_document_done.pdf`.

4.  **Using renamer on a directory with spaces in the name:**

    ```bash
    ./renamer -dir "/path/to/my files" -prefix "new_"
    ```
    If your directory contains spaces, enclose the directory path in quotation marks.

5. **Using renamer on a directory with special characters in the name:**
    ```bash
    ./renamer -dir "/path/to/my$files" -prefix "new_"
    ```
    If your directory contains special characters, enclose the directory path in quotation marks.

6. **Using renamer on a directory in the current working directory:**
    ```bash
    ./renamer -dir ./myfiles -prefix "new_"
    ```
    If your directory is in the current working directory, use `./` before the directory name.

## 🛠️ Installation 🛠️

1.  **Install Go:** ⬇️ Ensure you have Go installed on your system. You can download it from [golang.org](https://www.google.com/url?sa=E&source=gmail&q=https://golang.org/dl/). Follow the installation instructions for your operating system.

2.  **Clone the Repository:** 📥 Clone this repository to your local machine using Git. If you don't have Git installed, you can download it from [git-scm.com](https://git-scm.com/).

    ```bash
    git clone <repository_url>
    ```

    Replace `<repository_url>` with the actual URL of your Git repository.

3.  **Navigate to the Project Directory:** 🚶 Use the `cd` command to navigate to the project directory that you just cloned.

    ```bash
    cd renamer
    ```

4.  **Build the Executable:** 🔨 Run the following command to build the executable:

    ```bash
    go build -o renamer
    ```

    This will create an executable file named `renamer` (or `renamer.exe` on Windows) in the current directory. 🏁 The `-o renamer` part is crucial for correctly naming your executable.

    * **For Linux/macOS:** You might need to make the executable file executable:

        ```bash
        chmod +x renamer
        ```

    * **For Windows:** The `renamer.exe` file will be directly executable.

5. **Optional: Move the executable to a directory in your PATH**
    To make the `renamer` executable available from anywhere in your system, you can move it to a directory that is in your system's PATH environment variable.

    * **Linux/macOS:**
        ```bash
        sudo mv renamer /usr/local/bin/
        ```
    * **Windows:**
        Add the directory containing `renamer.exe` to your system's PATH environment variable through the system settings.

## 📦 Dependencies 📦

This project uses the standard Go library and does not require any external dependencies. 🎉 This keeps the tool lightweight and easy to use.

## 🤝 Contributing 🤝

Contributions are welcome! 🥳 If you find a bug 🐛, have suggestions for improvements 💡, or want to add new features 🚀, please feel free to contribute. Here's how:

1.  **Fork the repository:** Click the "Fork" button on the top right of the repository page.
2.  **Clone your fork:** Clone the forked repository to your local machine.

    ```bash
    git clone [https://github.com/your-username/renamer.git](https://www.google.com/search?q=https://github.com/your-username/renamer.git)
    ```

3.  **Create a new branch:** Create a branch for your changes.

    ```bash
    git checkout -b feature/your-feature-name
    ```

4.  **Make your changes:** Implement your changes or fixes.
5.  **Commit your changes:** Commit your changes with a descriptive message.

    ```bash
    git commit -m "Add your feature or fix"
    ```

6.  **Push your changes:** Push the changes to your forked repository.

    ```bash
    git push origin feature/your-feature-name
    ```

7.  **Create a pull request:** Go to the original repository on GitHub and create a pull request from your branch.

Please ensure your code adheres to the project's coding style and includes appropriate tests.

## 📜 License 📜

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT). ⚖️