# Template golang app

Шаблон для использования в проектах написанных на Go. Проект поддерживается структуры [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

## Go modules 

Для изменения имени модуля для go есть скрипты [tools](tools/). Скрипты сделаны если текстовый редактор не поддерживает возможности самостоятельно заменить имена и импорты.
```bash
$ ./tools/generate.sh <new name module>
```
Или же 
```bash
$ python ./tools/generate.py <new name module>
```