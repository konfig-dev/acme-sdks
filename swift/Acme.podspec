Pod::Spec.new do |s|
  s.name = 'Acme'
  s.swift_version = '5.0'
  s.ios.deployment_target = '9.0'
  s.osx.deployment_target = '10.11'
  s.tvos.deployment_target = '9.0'
  s.watchos.deployment_target = '3.0'
  s.version = '1.0.0'
  s.source = { :git => 'https://github.com/konfig-dev/acme-sdks.git', :tag => 'v1.0.0-swift' }
  s.authors = 'acme.com'
  s.license = 'MIT'
  s.homepage = 'https://github.com/konfig-dev/acme-sdks/tree/main/swift'
  s.summary = 'Acme Swift SDK'
  s.source_files = 'swift/Acme/**/*.swift'
  s.dependency 'AnyCodable-FlightSchool', '~> 0.6.1'
end
