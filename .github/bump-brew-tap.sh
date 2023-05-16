#!/usr/bin/env bash

set -e -x -u

shaOS=$(shasum -a 256 release/leftovers-*-darwin-amd64 | cut -d ' ' -f 1)
shaLinux=$(shasum -a 256 release/leftovers-*-linux-amd64 | cut -d ' ' -f 1)

pushd homebrew-tap
  cat <<EOF > leftovers.rb
class Leftovers < Formula
  desc "Command line utility for cleaning orphaned IAAS resources."
  homepage "https://github.com/genevieve/leftovers"
  version "${RELEASE_VERSION}"

  if OS.mac?
    url "https://github.com/genevieve/leftovers/releases/download/#{version}/leftovers-#{version}-darwin-amd64"
    sha256 "${shaOS}"
  elsif OS.linux?
    url "https://github.com/genevieve/leftovers/releases/download/#{version}/leftovers-#{version}-linux-amd64"
    sha256 "${shaLinux}"
  end

  depends_on :arch => :x86_64

  def install
    binary_name = "leftovers"
    if OS.mac?
      bin.install "leftovers-#{version}-darwin-amd64" => binary_name
    elsif OS.linux?
      bin.install "leftovers-#{version}-linux-amd64" => binary_name
    end
  end

  test do
    system "#{bin}/#{binary_name} --help"
  end
end
EOF

  cat leftovers.rb

  git add leftovers.rb
  if ! [ -z "$(git status --porcelain)" ];
  then
    git config --global user.email "leftovers-ci"
    git config --global user.name "leftovers-ci"
    git commit -m "Release leftovers ${RELEASE_VERSION}"
    git push
  else
    echo "No new version to commit"
  fi
popd
