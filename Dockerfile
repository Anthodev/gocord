ARG GO_VERSION=1.18
ARG CADDY_VERSION=2

FROM golang:${GO_VERSION}
ARG USERNAME=dev
ARG USER_UID=1000
ARG USER_GID=1000

RUN apt update \
    && apt upgrade -y \
    && apt install -y git bash

# Setup user
RUN addgroup --gid ${USER_GID} ${USERNAME}
RUN adduser $USERNAME --home /home/$USERNAME --disabled-password --uid $USER_UID --gid $USER_GID && \
    mkdir -p /etc/sudoers.d && \
    echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME && \
    chmod 0440 /etc/sudoers.d/$USERNAME

RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.2/zsh-in-docker.sh)" -- \
	-a 'HIST_STAMPS="yyyy-mm-dd"' \
	-p git \
	-p z \
	-p autojump \
	-p history \
	-p last-working-dir \
	-p docker \
	-p github \
	-p golang \
	-p https://github.com/zsh-users/zsh-autosuggestions \
	-p https://github.com/zsh-users/zsh-completions \
	-p https://github.com/zsh-users/zsh-syntax-highlighting \
	-p https://github.com/zsh-users/zsh-history-substring-search

RUN cp -af /root/.oh-my-zsh /home/${USERNAME}/ && \
  cp -af /root/.zshrc /home/${USERNAME}/ && sed -i 's/root/home\/${USERNAME}/g' /home/${USERNAME}/.zshrc && \
  chown -R ${USER_UID}:${USER_GID} /home/${USERNAME}/.oh-my-zsh && \
  chown -R ${USER_UID}:${USER_GID} /home/${USERNAME}/.zshrc

# Setup shell
USER $USERNAME

WORKDIR /usr/app

COPY . .

RUN go install github.com/cosmtrek/air@v1.29.0

#CMD ["air"]