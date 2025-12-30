# Expensia - expense tracker CLI application

## Overview
This is a simple command-line application designed to help users manage their 
personal finances. The application allows users to add, update, delete, and view
expenses. It also provides summaries of expenses, both overall and by month, 
making it easier to track spending habits and stay within budget.

## Features
- Add an expense with a description and amount  
- Update an existing expense  
- Delete an expense  
- View all expenses  
- View a summary of all expenses  
- View a summary of expenses for a specific month (current year)  

## Additional features
- Export expenses to JSON or CSV files 
- Add expense categories and filter expenses by category  
- Set a monthly budget and receive a warning when the budget is exceeded  

## Getting started

- Build
```bash
git clone https://github.com/Mirsait/expensia
cd expensia
go build -o expensia
```
- Installation
```bash
git clone https://github.com/Mirsait/expensia
cd expensia
sudo make install
```
- Uninstall
```bash
sudo make uninstall
```

## Using
```bash
$ expensia add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)

$ expensia add --description "Dinner" --amount 10
# Expense added successfully (ID: 2)

$ expensia list
# ID  Date       Description  Amount
# 1   2024-08-06  Lunch        $20
# 2   2024-08-06  Dinner       $10

$ expensia summary
# Total expenses: $30

$ expensia delete --id 2
# Expense deleted successfully

$ expensia summary
# Total expenses: $20

$ expensia summary --month 8
# Total expenses for August: $20
```

## License
[MIT License](LICENSE) — feel free to use, modify, and distribute.
