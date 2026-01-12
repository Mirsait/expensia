# Expensia - expense tracker CLI application

## Overview
This is a simple command-line application designed to help users manage their 
personal finances. The application allows users to add, update, delete, and view
expenses. It also provides summaries of expenses, both overall and by month, 
making it easier to track spending habits and stay within budget.

[roadmap.sh](https://roadmap.sh/projects/expense-tracker)

## Features
- Add an expense with a description, a category and amount  
- Update an existing expense  
- Delete an expense  
- View all expenses, expenses filtered by category  
- View a summary of all expenses  
- View a summary of expenses for a specific month (current year)  
- Export expenses to CSV file

## Getting started

- Build
```bash
git clone https://github.com/Mirsait/expensia
cd expensia
go build -o build/expensia
# or
make build
```
- Install
```bash
git clone https://github.com/Mirsait/expensia
cd expensia
make build
sudo make install
```
- Uninstall
```bash
sudo make uninstall
```

## Using
```bash
$ expensia add --description "Lunch" --amount 20 --category "food - lunch"
# Expense added successfully (ID: 1)

$ expensia add -D "Dinner" -a 10 -c "food - dinner"
# Expense added successfully (ID: 2)

$ expensia list
# ID        Date     Category  Description  Amount
# 1   2024-08-06   food-lunch        Lunch  $20
# 2   2024-08-06  food-dinner       Dinner  $10

$ expensia list -c "food-lunch"
# ID        Date     Category  Description  Amount
# 1   2024-08-06   food-lunch        Lunch  $20

$ expensia category
# food-lunch
# food-dinner

$ expensia summary
# Total expenses: $30

$ expensia delete --id 2
# expensia delete -i 2
# Expense deleted successfully

$ expensia summary
# Total expenses: $20

$ expensia summary --month 8
# expensia summary -m 8
# Total expenses for August: $20
```

## License
[MIT License](LICENSE) — feel free to use, modify, and distribute.
