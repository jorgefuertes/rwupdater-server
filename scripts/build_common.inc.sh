EXE_NAME="retroserver"

git status | egrep "On branch master|En la rama master" &> /dev/null
if [[ $? -ne 0 ]]
then
	echo "Refusing to build from other branch than master"
	exit 1
fi

git status | egrep "working tree clean|el árbol de trabajo está limpio" &> /dev/null
if [[ $? -ne 0 ]]
then
	echo "Commit local changes and create a new tag first"
	#exit 1
fi

git describe --tags --contains &> /dev/null
if [[ $? -eq 0 ]]
then
	VER=$(git describe --tags --contains)
else
	VER=$(git describe --tags)
	if [[ $? -ne 0 ]]
	then
		echo "Cannot get version tag, please check git status"
		exit 1
	fi
fi
echo "Version: ${VER}"

WHO=$(whoami)
TIME=$(date +"%d-%m-%Y@%H:%M:%S")
# darwin, linux, windows
# amd64, arm, arm64...
BUILD_LIST=(
	"darwin/amd64"
	"linux/amd64"
)

if [[ -f .build ]]
then
	BUILD=$(cat .build)
else
	BUILD=0
fi
BUILD=$(($BUILD + 1))
echo $BUILD > .build

FLAGS="-s -w \
	-X git.martianoids.com/queru/retroserver/internal/build.version=${VER} \
	-X git.martianoids.com/queru/retroserver/internal/build.user=${WHO} \
	-X git.martianoids.com/queru/retroserver/internal/build.time=${TIME} \
	-X git.martianoids.com/queru/retroserver/internal/build.number=${BUILD} \
"
