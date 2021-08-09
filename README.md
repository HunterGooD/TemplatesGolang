# Template golang app

Шаблон для использования в проектах написанных на Go. Проект поддерживается структуры [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

## Go modules 

Для изменения имени модуля для go есть скрипты [scripts](scripts/). Скрипты сделаны если текстовый редактор не поддерживает возможности самостоятельно заменить имена и импорты.
```bash
$ ./scripts/generate.sh <new name module>
```
Или же 
```bash
$ python ./scripts/generate.py <new name module>
```