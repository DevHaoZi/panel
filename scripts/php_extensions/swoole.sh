#!/bin/bash
export PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:$PATH

: '
Copyright (C) 2022 - now  HaoZi Technology Co., Ltd.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
'

HR="+----------------------------------------------------"

downloadUrl="https://git.haozi.net/opensource/download/-/raw/main/panel/php_extensions"
action="$1"
phpVersion="$2"
swooleVersion="5.1.2"

Install() {
    # 检查是否已经安装
    isInstall=$(cat /www/server/php/${phpVersion}/etc/php.ini | grep '^extension=swoole')
    if [ "${isInstall}" != "" ]; then
        echo -e $HR
        echo "PHP-${phpVersion} 已安装 swoole"
        exit 1
    fi

    cd /www/server/php/${phpVersion}/src/ext
    rm -rf swoole
    rm -rf swoole-src-${swooleVersion}.zip
    wget -T 60 -t 3 -O swoole-src-${swooleVersion}.zip ${downloadUrl}/swoole-src-${swooleVersion}.zip
    wget -T 20 -t 3 -O swoole-src-${swooleVersion}.zip.checksum.txt ${downloadUrl}/swoole-src-${swooleVersion}.zip.checksum.txt

    if ! sha256sum --status -c swoole-src-${swooleVersion}.zip.checksum.txt; then
        echo -e $HR
        echo "错误：PHP-${phpVersion} swoole 源码 checksum 校验失败，文件可能被篡改或不完整，已终止操作"
        exit 1
    fi

    unzip swoole-src-${swooleVersion}.zip
    mv swoole-src-${swooleVersion} swoole
    rm -f swoole-src-${swooleVersion}.zip
    rm -f swoole-src-${swooleVersion}.zip.checksum.txt
    cd swoole
    /www/server/php/${phpVersion}/bin/phpize
    ./configure --with-php-config=/www/server/php/${phpVersion}/bin/php-config
    make
    if [ "$?" != "0" ]; then
        echo -e $HR
        echo "PHP-${phpVersion} swoole 编译失败"
        exit 1
    fi
    make install
    if [ "$?" != "0" ]; then
        echo -e $HR
        echo "PHP-${phpVersion} swoole 安装失败"
        exit 1
    fi

    sed -i '/;haozi/a\extension=swoole' /www/server/php/${phpVersion}/etc/php.ini

    # 重载PHP
    systemctl reload php-fpm-${phpVersion}.service
    echo -e $HR
    echo "PHP-${phpVersion} swoole 安装成功"
}

Uninstall() {
    # 检查是否已经安装
    isInstall=$(cat /www/server/php/${phpVersion}/etc/php.ini | grep '^extension=swoole$')
    if [ "${isInstall}" == "" ]; then
        echo -e $HR
        echo "PHP-${phpVersion} 未安装 swoole"
        exit 1
    fi

    sed -i '/extension=swoole/d' /www/server/php/${phpVersion}/etc/php.ini

    # 重载PHP
    systemctl reload php-fpm-${phpVersion}.service
    echo -e $HR
    echo "PHP-${phpVersion} swoole 卸载成功"
}

if [ "$action" == 'install' ]; then
    Install
fi
if [ "$action" == 'uninstall' ]; then
    Uninstall
fi
